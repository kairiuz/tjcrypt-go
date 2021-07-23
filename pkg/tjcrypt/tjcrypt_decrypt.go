package tjcrypt

import (
	"encoding/binary"
	"errors"

	lz4 "github.com/patelh/golz4"
	"github.com/xxtea/xxtea-go/xxtea"
)

func decryptXxteaLZ4(in []byte) ([]byte, error) {
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

func decryptLZ4(in []byte) ([]byte, error) {
	finalSz := binary.LittleEndian.Uint32(in[3:7])
	uncompressedSz := binary.LittleEndian.Uint32(in[7:11])
	if uncompressedSz != finalSz {
		return nil, errors.New("decrypt mode z: finalSz != uncompressedSz")
	}

	compressed := in[11:]
	final := make([]byte, finalSz)
	_, err := lz4.Uncompress(compressed, final)
	if err != nil {
		return nil, err
	}

	return final, nil
}

// Decrypt tjcrypt buffer
func Decrypt(in []byte) ([]byte, error) {
	var (
		final []byte
		err   error
	)

	if !checkHeader(in) {
		return nil, errors.New("invalid header")
	}

	mode := in[2]
	switch mode {
	case '!':
		final, err = decryptXxteaLZ4(in)
	case 'z':
		final, err = decryptLZ4(in)
	case 'e':
		final, err = in[7:], nil // is this even an encryption
	}

	return final, err
}
