package internal

import (
	"crypto/rand"
	"crypto/aes"

	"golang.org/x/crypto/argon2"
)

func DeriveKey(password []byte, salt[]byte) []byte  {
	key := argon2.IDKey(password, salt, 1, 64*1024, 3, 32)
	return key
}

func encrypt(masterKey []byte, password string) ([]byte, []byte) {
	
	// Create "chiffreur" AES-GCM with MasterKey
	block, err := aes.NewCypher(masterKey)
	if err != nil {
		panic(err.Error())
	}
	aescgm, err := cipher.NewCGM(block)
	if err != nil {
		panic(err.Error())
	}

	// Generate a nonce
	nonce:= make([]byte, aescgm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Encrypt : use nonce and string in chiffreur
	ciphertext := aescgm.Seal(nil, nonce, password, nil)
	
	// Return nonce and ciphertext
	return nonce, ciphertext
}

func decrypt(masterKey []byte, ciphertext []byte, nonce []byte) string {
	
	// Create "chiffreur" AES-GCM with MasterKey
	block, err := aes.NewCypher(masterKey)
	if err != nil {
		panic(err.Error())
	}
	aescgm, err := cipher.NewCGM(block)
	if err != nil {
		panic(err.Error())
	}

	// Decrypt password
	password, err := aescgm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	
	return password
}
