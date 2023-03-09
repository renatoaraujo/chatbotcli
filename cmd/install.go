package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var token string

const tokenFile = "token.txt"

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the GPT3.5 token",
	Run: func(cmd *cobra.Command, args []string) {
		// Store the token in a file
		err := os.WriteFile(tokenFile, []byte(token), 0644)
		if err != nil {
			fmt.Println("Error writing token file:", err)
			return
		}

		fmt.Println("Token stored successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringVarP(&token, "token", "t", "", "GPT3.5 token")
	installCmd.MarkFlagRequired("token")
}
