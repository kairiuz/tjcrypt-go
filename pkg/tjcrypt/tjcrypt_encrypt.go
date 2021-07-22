package tjcrypt

import (
	"encoding/binary"

	lz4 "github.com/patelh/golz4"
	"github.com/xxtea/xxtea-go/xxtea"
)

// Encrypt tjcrypt buffer
func Encrypt(in []byte) ([]byte, error) {
	// compress
	commpressSz := lz4.CompressBound(in)
	compressed := make([]byte, commpressSz+4)
	if _, err := lz4.Compress(in, compressed[4:]); err != nil {
		return nil, err
	}

	// put dstSz
	dstSz := uint32(len(in))
	binary.LittleEndian.PutUint32(compressed, dstSz)

	// prepare key
	fixKey := make([]byte, 16)
	for i := 0; i < 16; i++ {
		fixKey[i] = TJ_DEFAULT_KEY[i%16] ^ TJ_PASSWORD[i]
	}

	encrypted := xxtea.Encrypt(compressed, fixKey)

	// dstSz
	bufSz := make([]byte, 4)
	binary.LittleEndian.PutUint32(bufSz, dstSz)

	out := []byte{'t', 'j', '!'}
	out = append(out, []byte(TJ_PASSWORD)...)
	out = append(out, bufSz...)
	out = append(out, encrypted...)

	return out, nil
}
