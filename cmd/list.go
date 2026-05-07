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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored account names.",
	Long: "List all stored account names.",
	Run: func(cmd *cobra.Command, args []string) {

		// Read file
		content, err := ioutil.ReadFile("vault.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		// Extract payload
		var payload internal.Vault
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
		
		// Print list
		for i := 0; i < len(payload.Entries); i++ {
			fmt.Println(payload.Entries[i].Account)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
