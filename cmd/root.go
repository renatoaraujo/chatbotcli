package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatbotcli",
	Short: "A CLI tool for interacting with GPT3.5 chatbots",
	Long: `chatbotcli is a command-line interface (CLI) tool for interacting with chatbots powered by GPT3.5.
It allows you to quickly and easily generate text responses to your queries.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Welcome to chatbotcli!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Initialize configuration here
}
