/*
Copyright © 2022 Basin authors@basin.dev
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Make changes to existing metadata",
	Long: `Make changes to existing metadata, including
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
	rootCmd.AddCommand(modifyCmd)
	modifyCmd.PersistentFlags().StringP("mode", "m", "", "consumer or producer mode")
	modifyCmd.MarkPersistentFlagRequired("mode")
	// https://github.com/spf13/cobra/blob/main/user_guide.md#flag-groups
	// rootCmd.MarkFlagsRequiredTogether("username", "password")
}
