package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask the OpenAI GPT-3.5 a question",
	Run: func(cmd *cobra.Command, args []string) {
		// Create a temporary file to store the user input
		file, err := os.CreateTemp("", "input")
		if err != nil {
			fmt.Println(err)
			return
		}

		// Open the file in vim for the user to input their question
		cmdVim := exec.Command("vim", "-c", "startinsert", file.Name())
		cmdVim.Stdin = os.Stdin
		cmdVim.Stdout = os.Stdout
		cmdVim.Stderr = os.Stderr
		err = cmdVim.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Read the user input from the file
		content, err := os.ReadFile(file.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		// Read the GPT3.5 token from the file
		tokenBytes, err := os.ReadFile(tokenFile)
		if err != nil {
			fmt.Println("Error reading token file:", err)
			return
		}
		token := string(tokenBytes)

		// Initialize the OpenAI client with the GPT3.5 token
		client := openai.NewClient(token)

		// Send the user input to the chatbot and print the response
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: string(content),
					},
				},
			},
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(resp.Choices[0].Message.Content)
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
