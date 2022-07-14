/*
Copyright © 2022 Basin authors@basin.dev
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "View current metadata",
	Long: `View current metadata, including
	- data
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
				fmt.Println("check in consumer mode called with sources entity")
			default:
				fmt.Println("error: check in consumer mode must be called with sources entity")
			}
		case "producer":
			switch entity {
			case "sources":
				fmt.Println("check in producer mode called with sources entity")
			default:
				fmt.Println("error: check in producer mode must be called with sources entity")
			}
		default:
			fmt.Println("error: checks must be called in either consumer or producer mode with sources entity.")
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.PersistentFlags().StringP("mode", "m", "", "consumer or producer mode")
	checkCmd.MarkPersistentFlagRequired("mode")
	checkCmd.PersistentFlags().StringP("entity", "e", "", "entities (e.g. sources)")
	checkCmd.MarkPersistentFlagRequired("entity")
	// https://github.com/spf13/cobra/blob/main/user_guide.md#flag-groups
	// rootCmd.MarkFlagsRequiredTogether("username", "password")
}
