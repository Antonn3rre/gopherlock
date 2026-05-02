/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"syscall"
	"log"
	"crypto/rand"
	"crypto/sha256"
	"gopherlock/internal"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// 1. Asks for password
		fmt.Println("Initialisation...")

		masterPassword := ""

		for {
		fmt.Println("Please provide your master key: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        log.Fatal(err)
    }
    password := strings.TrimSpace(string(bytePassword))
		if password == "" {
			continue
		}

		fmt.Println("\nPlease provide your master key: (again)")
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
	salt:= make([]byte, 16)
	rand.Read(salt)
	// TODO: check error

	// Create Master Key : (masterPassword + salt --> Argon2id)
	masterKey := internal.DeriveKey([]byte(masterPassword), []byte(salt))
	// Hash masterKey
	hashedMasterKey := sha256.Sum256(masterKey)

	// 3. Create the vault file
	v:= internal.Vault{ Salt: salt, CheckHash: hashedMasterKey[:], Entries: []internal.Entry{} }
	data, _ := json.MarshalIndent(v, "", " ")
	os.WriteFile("vault.json", data, 0600)
	},
}

//func init() -> enregistre la commande aupres de la racine + configure les options
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
