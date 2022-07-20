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
	- sources
	- permissions
	- royalties
	- schemas
	- cache
	- wallet
	as a consumer or producer.`,
	Run: func(cmd *cobra.Command, args []string) {

		mode, _ := cmd.Flags().GetString("mode")
		entity, _ := cmd.Flags().GetString("entity")

		switch mode {
		case "consumer":
			switch entity {
			case "sources":
				fmt.Println("list in consumer mode called with sources entity")
			default:
				fmt.Println("error: list in consumer mode must be called with sources entity")
			}
		case "producer":
			switch entity {
			case "sources":
				fmt.Println("list in producer mode called with sources entity")
			default:
				fmt.Println("error: list in producer mode must be called with sources entity")
			}
		default:
			fmt.Println("error: list must be called in either consumer or producer mode with sources entity.")
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("mode", "m", "", "consumer or producer mode")
	listCmd.MarkPersistentFlagRequired("mode")
	listCmd.PersistentFlags().StringP("entity", "e", "", "entities (e.g. sources)")
	listCmd.MarkPersistentFlagRequired("entity")
	// https://github.com/spf13/cobra/blob/main/user_guide.md#flag-groups
	// rootCmd.MarkFlagsRequiredTogether("username", "password")
}
