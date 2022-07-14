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
		entity, _ := cmd.Flags().GetString("entity")

		switch mode {
		case "consumer":
			switch entity {
			case "sources":
				fmt.Println("modify in consumer mode called with sources entity")
			default:
				fmt.Println("error: modify in consumer mode must be called with sources entity")
			}
		case "producer":
			switch entity {
			case "sources":
				fmt.Println("modify in producer mode called with sources entity")
			default:
				fmt.Println("error: modify in producer mode must be called with sources entity")
			}
		default:
			fmt.Println("error: modify must be called in either consumer or producer mode with sources entity.")
		}

	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)
	modifyCmd.PersistentFlags().StringP("mode", "m", "", "consumer or producer mode")
	modifyCmd.MarkPersistentFlagRequired("mode")
	modifyCmd.PersistentFlags().StringP("entity", "e", "", "entities (e.g. sources)")
	modifyCmd.MarkPersistentFlagRequired("entity")
	// https://github.com/spf13/cobra/blob/main/user_guide.md#flag-groups
	// rootCmd.MarkFlagsRequiredTogether("username", "password")
}
