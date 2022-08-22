/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sestinj/basin-node/client"
	"github.com/spf13/cobra"
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Request subscription to a resource",
	Long: `Request subscription to a resource. Input the necessary metadata, including
	- permissions`,
	Run: func(cmd *cobra.Command, args []string) {
		action, err := cmd.Flags().GetString("action") // TODO[FEATURE][1]: Should be able to specify multiple actions like `--action read --action write`.
		url := args[0]
		did := args[1]
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse flags: %s", err.Error())
		}

		// capabilities, err := cmd.Flags().GetString("capabilities")
		// cdata, err := util.UnmarshalFromFile[[]client.CapabilitySchema](capabilities)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "Could not read file: %s", err.Error())
		// }

		// TODO[FEATURE][1]: flag for expiration
		time := time.Now().Add(time.Hour * 24 * 365 * 10)
		// TODO[FEATURE][1]: Should also be able to just specify a permissions file

		permission := client.PermissionJson{Entities: []string{did}, Data: []string{url}, Capabilities: []client.CapabilitySchema{{Action: &action, Expiration: &time}}}

		ctx := context.Background()

		subscribeRequest := client.NewSubscribeRequest([]client.PermissionJson{permission})

		body, r, err := interactiveConfig.ApiClient.DefaultApi.Subscribe(ctx).SubscribeRequest(*subscribeRequest).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to subscribe to resource: %s", err.Error())
		} else if r.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "Response returned error: %s", r.Status)
		}
		fmt.Fprintf(os.Stdout, "The subscription request has been noted: %s\n", body)
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)
	// subscribeCmd.Flags().StringP("capabilities", "c", "", "Path to capabilities file")
	// subscribeCmd.MarkFlagRequired("capabilities")
	subscribeCmd.Flags().StringP("action", "a", "", "Action requesting permissions for (read/write)")
	subscribeCmd.MarkFlagRequired("action")
}
