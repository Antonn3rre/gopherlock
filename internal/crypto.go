package internal

import (
	"golang.org/x/crypto/argon2"
)

func DeriveKey(password []byte, salt[]byte) []byte  {
	key := argon2.IDKey(password, salt, 1, 64*1024, 3, 32)
	return key
}

// func encrypt
// func decrypt
