package tjcrypt

import (
	"encoding/binary"
	"errors"

	lz4 "github.com/patelh/golz4"
	"github.com/xxtea/xxtea-go/xxtea"
)

// Encrypt into tjcrypt buffer with provided key
func EncryptWithCustomKey(in, key []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("key length must be 16")
	}

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
		fixKey[i] = TJ_DEFAULT_KEY[i%16] ^ key[i]
	}

	encrypted := xxtea.Encrypt(compressed, fixKey)

	// dstSz
	bufSz := make([]byte, 4)
	binary.LittleEndian.PutUint32(bufSz, dstSz)

	out := []byte{'t', 'j', '!'}
	out = append(out, []byte(key)...)
	out = append(out, bufSz...)
	out = append(out, encrypted...)

	return out, nil
}

// Encrypt tjcrypt buffer
func Encrypt(in []byte) ([]byte, error) {
	return EncryptWithCustomKey(in, []byte(TJ_PASSWORD))
}
