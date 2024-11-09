/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

func generateRandomPassword(length int, sC bool) string {
	alphanumericals := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialChars := "!@#$%^&*()_+-=[]{}|;:,.<>?~`'"
	allChars := ""
	result := make([]byte, length)

	if sC {
		allChars = alphanumericals + specialChars
	} else {
		allChars = alphanumericals
	}

	for i := 0; i < length; i++ {
		result[i] = allChars[rand.Intn(len(allChars))]
	}

	return string(result)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Explain passwordGen",
	Short: `PasswordGen is a CLI application for password generation.`,
	Long: `PasswordGen is a CLI application for password generation.

  --length, -t    to specify password length. (default 8)
  --special, -s   to include special characters in the generated password`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		sC, _ := cmd.Flags().GetBool("special")

		generatedPassword := generateRandomPassword(length, sC)
		fmt.Println("Generated Password: ", generatedPassword)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.passwordGen-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("special", "s", false, "--spacial/-s flag is used to include special characters in the generated password")
	rootCmd.Flags().IntP("length", "l", 8, "--length/-t is used to   to specify password length (default 8)")
}
