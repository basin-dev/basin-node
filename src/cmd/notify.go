/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sestinj/basin-node/client"
)

// notifyCmd represents the notify command
// basin notify read <URL>
var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Notify the network of new data at a resource",
	Long:  `Notify the network of new data at a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		if !interactive {
			fmt.Fprintln(os.Stderr, "This command can only be run in interactive mode. Use `basin attach` first.")
		}
		url := args[0]

		ctx := context.Background()

		notifyReq := client.NewNotifyRequest(url)
		resp, r, err := interactiveConfig.ApiClient.DefaultApi.Notify(ctx).NotifyRequest(*notifyReq).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to notify: %s", err.Error())
		} else if r.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "Failed to notify: %s", r.Status)
		}
		fmt.Fprintln(os.Stdout, resp)
	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)
}
