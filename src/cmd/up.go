/*
Copyright © 2022 Basin authors@basin.dev
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Starts node.",
	Long:  `EXPLAIN WHAT BRINGING A NODE UP DOES.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("up called")
		detach, err := cmd.Flags().GetBool("detach")
		if detach {
			fmt.Println("detached")
		}
		fmt.Println(err)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	upCmd.Flags().BoolP("detach", "d", false, "Detached mode: run node in the background")
}
