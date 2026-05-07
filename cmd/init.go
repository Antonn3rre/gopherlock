/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"gopherlock/internal"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new vault with a master password. This creates the `vault.json` file which stores all your credentials securely.",
	Long: "Initialize a new vault with a master password. This creates the `vault.json` file which stores all your credentials securely.",
	Run: func(cmd *cobra.Command, args []string) {

		// 1. Asks for password
		fmt.Println("Initialisation...")

		masterPassword := ""

		for {
			fmt.Print("Please provide your master password: ")
			bytePassword, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatal(err)
			}
			password := strings.TrimSpace(string(bytePassword))
			if password == "" {
				continue
			}

			fmt.Print("\nPlease provide your master password (again): ")
			bytePassword, err = term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatal(err)
			}
			copiedPassword := strings.TrimSpace(string(bytePassword))
			if copiedPassword != password {
				fmt.Println("Error: the passwords do not match")
				continue
			}
			masterPassword = password
			break
		}
		fmt.Printf("Your password is: |%s|\n", masterPassword)

		// 2. Create salt
		salt := make([]byte, 16)
		_, err := rand.Read(salt)
		if err != nil {
			panic(err.Error())
		}

		// Create Master Key : (masterPassword + salt --> Argon2id)
		masterKey := internal.DeriveKey([]byte(masterPassword), []byte(salt))
		// Hash masterKey
		hashedMasterKey := sha256.Sum256(masterKey)

		// 3. Create the vault file
		v := internal.Vault{Salt: salt, CheckHash: hashedMasterKey[:], Entries: []internal.Entry{}}
		data, _ := json.MarshalIndent(v, "", " ")
		os.WriteFile("vault.json", data, 0600)
	},
}

// func init() -> enregistre la commande aupres de la racine + configure les options
func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
