package internal

import (
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"golang.org/x/term"
	"syscall"
	"crypto/sha256"
	"bytes"
)

func Login() []byte {

	// Retrieve salt and checkHash
	content, err := ioutil.ReadFile("vault.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Vault
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	salt := payload.Salt
	checkHash := payload.CheckHash

	// Ask for master password
	fmt.Print("\nPlease provide your master passwrd: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
  if err != nil {
  	log.Fatal(err)
  }

	// Create Key : (password + salt --> Argon2id)
	key := DeriveKey([]byte(password), []byte(salt))
	// Hash Key
	hashedKey := sha256.Sum256(key)

	// Compare checkHash to key
	if ! bytes.Equal(checkHash[:], hashedKey[:]) {
		log.Fatal("Erreur")
		// fatal error
	}

	fmt.Println("\nLogin successfull")
	return checkHash
}
