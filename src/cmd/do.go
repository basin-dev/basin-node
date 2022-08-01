/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/sestinj/basin-node/client"
)

// doCmd represents the do command
// basin do read <URL>
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Perform an action on a resource",
	Long: `Perform an action on a resource, for example:
	- read
	- write
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if !interactive {
			fmt.Fprintln(os.Stderr, "This command can only be run in interactive mode. Use `basin attach` first.")
		}
		action := args[0]
		url := args[1]

		cfg := client.NewConfiguration()
		apiClient := client.NewAPIClient(cfg)
		ctx := context.Background()

		switch action {
		case "read":
			resp, r, err := apiClient.DefaultApi.Read(ctx).Url(url).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read resource: %s", err.Error())
			} else if r.StatusCode != 200 {
				fmt.Fprintf(os.Stderr, "Failed to read resource: %s", r.Status)
			}
			fmt.Fprintln(os.Stdout, []byte(resp)) // TODO: Can we decode the bytes for the user? Where should this happen?????
		case "write":
			if len(args) < 3 {
				log.Fatal("Not enough arguments supplied to write command.")
			}
			value := args[2]
			writeReq := client.NewWriteRequest(url, value)
			resp, r, err := apiClient.DefaultApi.Write(ctx).WriteRequest(*writeReq).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read resource: %s", err.Error())
			} else if r.StatusCode != 200 {
				fmt.Fprintf(os.Stderr, "Failed to read resource: %s", r.Status)
			}
			fmt.Fprintln(os.Stdout, resp)
		default:
			log.Fatal("Arbitrary actions are not yet supported")
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
