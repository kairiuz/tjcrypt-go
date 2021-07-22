package tjcrypt

import (
	"bytes"
	"testing"
)

func TestEncryptDecypt(t *testing.T) {
	plain := []byte("-- just a simple stub")
	want := []byte("tj!MODDEDBYPL0NK!!!\x15\x00\x00\x00\xfa\xa6\xb5\xc1\xb9\x96-\xf0.\x9e\xc0\xf6\x98jغ\x8c\xf7$\x86\xd8'\x9arS\x14?\xe0pD?\xf5\xa6\vQ\x96Yƭ\xe74\xa3Ȳ\xec\xdbC\x99")

	enc, err := Encrypt(plain)
	if !bytes.Equal(enc, want) || err != nil {
		t.Fatalf(`Encrypt(plain) = %q, %v, want %#q`, enc, err, want)
	}

	dec, err := Decrypt(enc)
	if !bytes.Equal(dec, plain) || err != nil {
		t.Fatalf(`Decrypt(enc) = %q, %v, want %#q`, enc, err, plain)
	}
}
