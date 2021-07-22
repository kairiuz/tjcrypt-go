package tjcrypt

import (
	"encoding/binary"
	"errors"

	lz4 "github.com/patelh/golz4"
	"github.com/xxtea/xxtea-go/xxtea"
)

// Decrypt tjcrypt buffer
func Decrypt(in []byte) ([]byte, error) {
	if !checkHeader(in) {
		return nil, errors.New("invalid header")
	}

	// prepare key
	fixKey := make([]byte, 16)
	for i := 0; i < 16; i++ {
		fixKey[i] = TJ_DEFAULT_KEY[i%16] ^ in[i+3]
	}

	// decrypt
	decrypted := xxtea.Decrypt(in[23:], fixKey)

	// sanity check
	dstSz := binary.LittleEndian.Uint32(decrypted[:4])
	if dstSz != binary.LittleEndian.Uint32(in[19:23]) {
		return nil, errors.New("decrypted dstSz != encrypted dstSz")
	}

	// decompress
	compressed := decrypted[4:]
	final := make([]byte, dstSz)
	_, err := lz4.Uncompress(compressed, final)
	if err != nil {
		return nil, err
	}

	return final, nil
}
