/*
Copyright © 2022 Basin authors@basin.dev
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View current metadata",
	Long: `View current metadata, including
	- permissions
	- royalties
	- schemas
	- wallet
	as a consumer or producer.`,
	Run: func(cmd *cobra.Command, args []string) {

		op, _ := cmd.Flags().GetString("op")
		metadata, _ := cmd.Flags().GetString("metadata")

		method := "LevelDbHttpAdapter.Read(%s)\n"

		switch op {
		case "consumer":
			switch metadata {
			case "permissions":
				url := "basin://local.consumer.permissions.metadata"
				fmt.Printf(method, url)
			case "requests":
				url := "basin://local.consumer.requests.metadata"
				fmt.Printf(method, url)
			case "royalties":
				url := "basin://local.consumer.royalties.metadata"
				fmt.Printf(method, url)
			case "schemas":
				url := "basin://local.consumer.schemas.metadata"
				fmt.Printf(method, url)
			case "wallet":
				url := "basin://local.consumer.wallet.metadata"
				fmt.Printf(method, url)
			default:
				fmt.Println("error: list in consumer mode must be called with an entity")
			}
		case "producer":
			switch metadata {
			case "permissions":
				url := "basin://local.producer.permissions.metadata"
				fmt.Printf(method, url)
			case "requests":
				url := "basin://local.producer.requests.metadata"
				fmt.Printf(method, url)
			case "royalties":
				url := "basin://local.producer.royalties.metadata"
				fmt.Printf(method, url)
			case "schemas":
				url := "basin://local.producer.schemas.metadata"
				fmt.Printf(method, url)
			case "wallet":
				url := "basin://local.producer.wallet.metadata"
				fmt.Printf(method, url)
			default:
				fmt.Println("error: list in producer mode must be called with an entity")
			}
		default:
			fmt.Println("error: list must be called in either consumer or producer mode with an entity.")
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("op", "o", "", "consumer or producer mode")
	listCmd.MarkPersistentFlagRequired("op")
	listCmd.PersistentFlags().StringP("metadata", "m", "", "metadata (e.g. permissions)")
	listCmd.MarkPersistentFlagRequired("metadata")
	// https://github.com/spf13/cobra/blob/main/user_guide.md#flag-groups
	// rootCmd.MarkFlagsRequiredTogether("username", "password")
}
