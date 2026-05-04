package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/argon2"
)

func DeriveKey(password []byte, salt []byte) []byte {
	key := argon2.IDKey(password, salt, 1, 64*1024, 3, 32)
	return key
}

func Encrypt(masterKey []byte, password []byte) ([]byte, []byte) {

	// Create "chiffreur" AES-GCM with MasterKey
	block, err := aes.NewCipher(masterKey)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// Generate a nonce
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Encrypt : use nonce and string in chiffreur
	ciphertext := aesgcm.Seal(nil, nonce, password, nil)

	// Return nonce and ciphertext
	return nonce, ciphertext
}

func Decrypt(masterKey []byte, ciphertext []byte, nonce []byte) []byte {

	// Create "chiffreur" AES-GCM with MasterKey
	block, err := aes.NewCipher(masterKey)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// Decrypt password
	password, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return password
}
