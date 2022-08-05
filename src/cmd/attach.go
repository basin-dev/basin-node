/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func GetInput(prompt string, reader *bufio.Reader) string {
	fmt.Printf("%s ", prompt)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	return input
}

const PROMPT = "~> "

func RunConsole() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`
	______  ___   _____ _____ _   _ 
	| ___ \/ _ \ /  ___|_   _| \ | |
	| |_/ / /_\ \\ '--.  | | |  \| |
	| ___ \  _  | '--. \ | | | . ' |
	| |_/ / | | |/\__/ /_| |_| |\  |
	\____/\_| |_/\____/ \___/\_| \_/
	`)
	prompt := func() string {
		return GetInput(PROMPT, reader)
	}
	for {
		input := prompt()
		// TODO: How to manually pass the input into Cobra CLI?

		fmt.Println(input)
	}
}

// attachCmd represents the attach command
// basin attach read <URL>
var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a Basin node and enter interactive CLI",
	Long:  `Attach to a Basin node and enter interactive CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		httpUrl, _ := cmd.Flags().GetString("http")

		interactive = true
		interactiveConfig.Url = httpUrl

		RunConsole()

		interactive = false
	},
}

func init() {
	rootCmd.AddCommand(attachCmd)
	attachCmd.Flags().String("http", "http://127.0.0.1:8555", "The URL where the node to connect to is being served.")
	attachCmd.MarkFlagRequired("http")
}
