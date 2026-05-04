/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gopherlock/internal"
	"io/ioutil"
	"log"
	"encoding/json"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")

		// check args
		if len(args) == 0 {
			fmt.Println("Missing arguments: please provide the account")
			return
		}
		account := args[0]

		// login
		hashedMaster := internal.Login()

		// Read vault
		content, err := ioutil.ReadFile("vault.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		var payload internal.Vault
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		found := false
		for i := 0; i < len(payload.Entries); i++ {
			if (payload.Entries[i].Account == account) {
				fmt.Println("")
				found = true
				password := internal.Decrypt(hashedMaster, payload.Entries[i].Ciphertext, payload.Entries[i].Nonce)
				fmt.Println("Username: ", payload.Entries[i].Username)
				fmt.Println("Password: ", string(password))
			}
		}
		if (!found) {
			fmt.Println("There's no password linked to this account")
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
