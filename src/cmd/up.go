/*
Copyright © 2022 Basin authors@basin.dev
*/
package cmd

import (
	"context"
	"log"

	"github.com/sestinj/basin-node/node"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start a Basin node",
	Long:  `Start a Basin node`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("UP")

		httpUrl, err := cmd.Flags().GetString("http")
		if err != nil { // TODO: @TyDunn, does the value argument to the flag creation functions below set the default value? Why would an error occur?
			httpUrl = ""
		}

		did, err := cmd.Flags().GetString("did")
		if err != nil {
			did = ""
		}

		pw, err := cmd.Flags().GetString("pw")
		if err != nil {
			pw = ""
		}

		config := node.BasinNodeConfig{
			Http: httpUrl,
			Did:  did,
			Pw:   pw,
		}
		config.SetDefaults()
		node.StartEverything(context.Background(), config)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
	upCmd.Flags().String("http", "", "Set the URL from which to host the node's HTTP interface.")
	upCmd.Flags().String("did", "", "Set the default signing DID for the node.")
	upCmd.Flags().String("pw", "", "Enter the password to decrypt the keystore file for the given DID")
	upCmd.MarkFlagsRequiredTogether("did", "pw")
}
