package sm2

import (
	"crypto/rand"
	"encoding/pem"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func TestEnCrypt(t *testing.T) {

	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
		return
	}

	publicKey := &privateKey.PublicKey
	sm2PublicKey, err := x509.MarshalSm2PublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
		return
	}

	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: sm2PublicKey,
	}

	memory := pem.EncodeToMemory(pubKeyBlock)

	encrypt, err := Encrypt(string(memory), "hello world", true)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(encrypt)
	t.Log(string(encrypt))

}

func TestDecrypt(t *testing.T) {
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
		return
	}

	sm2PrivateKey, err := x509.MarshalSm2PrivateKey(privateKey, nil)
	if err != nil {
		t.Fatal(err)
		return
	}
	priKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: sm2PrivateKey,
	}

	memory := pem.EncodeToMemory(priKeyBlock)

	decrypt, err := Decrypt(string(memory), "", string([]byte{4, 55, 91, 100, 0, 5, 34, 216, 147, 75, 240, 110, 77, 245, 214, 27, 107, 74, 182, 111, 114, 91, 39, 51, 100, 213, 157, 161, 117, 101, 134, 63, 197, 218, 8, 10, 95, 172, 3, 12, 169, 37, 202, 41, 46, 66, 161, 155, 250, 144, 115, 143, 255, 102, 5, 212, 223, 80, 10, 46, 137, 21, 19, 23, 35, 43, 88, 175, 219, 85, 143, 140, 46, 98, 173, 171, 225, 123, 137, 202, 87, 132, 0, 148, 118, 202, 204, 113, 217, 103, 244, 135, 85, 43, 218, 170, 127, 124, 201, 168, 103, 129, 193, 132, 126, 105, 34, 176}), true)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(decrypt))

}
