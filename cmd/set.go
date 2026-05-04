/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gopherlock/internal"
	"fmt"
	"golang.org/x/term"
	"log"
	"syscall"
	"bufio"
	"os"
	"io/ioutil"
	"encoding/json"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")

		// Login
		hashedMaster := internal.Login()
		fmt.Println("Login successfull")

		// Init scanner to read input
		scanner := bufio.NewScanner(os.Stdin)
		var account, username string
		
		// Ask about the account info
		fmt.Println("Account name (ex: Gmail, Wiki): ")
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

		// TODO: Check doublon account + password

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
			Account: account,
			Username: username,
			Ciphertext: ciphertext,
			Nonce: nonce,
		}

		payload.Entries = append(payload.Entries, *newEntry)
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
