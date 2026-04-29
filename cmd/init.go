/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"syscall"
	"log"

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
