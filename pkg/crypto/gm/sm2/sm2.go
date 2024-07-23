package sm2

import (
	"crypto/rand"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func Encrypt(pubKey string, plainText string) ([]byte, error) {
	pem, err := x509.ReadPublicKeyFromPem([]byte(pubKey))
	encrypt, err := sm2.Encrypt(pem, []byte(plainText), rand.Reader, sm2.C1C3C2)

	return encrypt, err
}

func Decrypt(privKey string, cipherText []byte) ([]byte, error) {
	pem, err := x509.ReadPrivateKeyFromPem([]byte(privKey), nil)
	decrypt, err := sm2.Decrypt(pem, cipherText, sm2.C1C3C2)
	return decrypt, err
}
