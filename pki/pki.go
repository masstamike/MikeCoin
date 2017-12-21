package pki

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"crypto/sha1"
)

func GenerateKeyPair () (crypto.PrivateKey){
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return privateKey
}

func Sign (s string, key *rsa.PrivateKey) {
	hash := sha1.Sum([]byte(s))
	rsa.SignPKCS1v15(nil, key, crypto.SHA1, hash[:])
}

func ValidSignature (message string, cipherText string, pubKey crypto.PublicKey) bool {
	return ValidSignature(message, cipherText, pubKey)
}
