/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
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

	for {
		input := strings.Split(GetInput(PROMPT, reader), " ")

		// Special interactive mode commands
		switch input[0] {
		case "basin":
			fmt.Println("No need to specify the basin prefix when in interactive mode, but your command was still parsed.")
			input = input[1:]
		case "quit", "exit":
			fmt.Println("Exiting session...")
			return
		case "up":
			fmt.Println("Cannot start node from within interactive terminal.")
		}

		// basin register basin://ty.com.twitter --adapter localhost:/8555 --schema=schema.json --permissions=permissions.json
		// go run . register basin://ty.com.twitter --adapter localhost:/8555 --schema=schema.json --permissions=permissions.json
		// register basin://tydunn.com.twitter.followers -a ../testing/config/adapter.json -p ../testing/config/permissions.yaml -s ../testing/config/schema.json
		// do read basin://tydunn.com.twitter.followers

		command, args, err := rootCmd.Find(input)
		if err != nil {
			log.Printf("Unknown Command to execute : %s\n", input)
			continue
		}

		err = command.ParseFlags(args)
		if err != nil {
			log.Printf("Err parsing flags: %s\n", err.Error())
			continue
		}

		command.Run(command, args)
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
