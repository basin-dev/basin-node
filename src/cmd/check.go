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

		if mode == "consumer" {
			fmt.Println("modify called in consumer mode.")
		} else if mode == "producer" {
			fmt.Println("modify called in producer mode.")
		} else {
			fmt.Println("modify must be called in either consumer or producer mode.")
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.PersistentFlags().StringP("mode", "m", "", "consumer or producer mode")
	checkCmd.MarkPersistentFlagRequired("mode")
	// https://github.com/spf13/cobra/blob/main/user_guide.md#flag-groups
	// rootCmd.MarkFlagsRequiredTogether("username", "password")
}
