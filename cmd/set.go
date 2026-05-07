/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/term"
	"gopherlock/internal"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Add a new password entry to the vault.",
	Long: "Add a new password entry to the vault.",
	Run: func(cmd *cobra.Command, args []string) {

		// Login
		hashedMaster := internal.Login()

		// Init scanner to read input
		scanner := bufio.NewScanner(os.Stdin)
		var account, username string

		// Ask about the account info
		fmt.Print("Account name (ex: Gmail, Wiki): ")
		if scanner.Scan() {
			account = scanner.Text()
		}

		// Username
		fmt.Println("Username: ")
		if scanner.Scan() {
			username = scanner.Text()
		}

		// Password
		fmt.Println("Password: ")
		password, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}

		// encrypt
		nonce, ciphertext := internal.Encrypt(hashedMaster, password)

		// Write in file
		content, err := ioutil.ReadFile("vault.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		var payload internal.Vault
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		newEntry := &internal.Entry{
			Account:    account,
			Username:   username,
			Ciphertext: ciphertext,
			Nonce:      nonce,
		}

		existingIndex := -1
		for i, entry := range payload.Entries {
			if entry.Account == account && entry.Username == username {
				existingIndex = i
				break
			}
		}

		if existingIndex != -1 {
			fmt.Print("A password already exists for this account and username, update it? (y/n): ")
			var confirm string
			if scanner.Scan() {
				confirm = strings.TrimSpace(strings.ToLower(scanner.Text()))
			}

			if confirm != "y" {
				fmt.Println("Cancelled.")
				return
			}

			payload.Entries[existingIndex] = *newEntry
		} else {
			payload.Entries = append(payload.Entries, *newEntry)
		}

		data, _ := json.MarshalIndent(payload, "", " ")
		os.WriteFile("vault.json", data, 0600)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
