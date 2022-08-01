/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new resource",
	Long: `Register a new resource as a producer. Input the necessary metadata, including
	- adapter
	- permissions
	- schema`,
	Run: func(cmd *cobra.Command, args []string) {
		adapter, err := cmd.Flags().GetString("adapter")
		permissions, err = cmd.Flags().GetString("permissions")
		schema, err = cmd.Flags().GetString("schema")
		url := args[0]
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse flags: %s", err.Error())
		}

		err = Register(context.Background(), url, adapter, permissions, schema)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to register resource: %s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("adapter", "a", "", "Path to adapter file")
	registerCmd.Flags().StringP("permissions", "p", "", "Path to permissions file")
	registerCmd.Flags().StringP("schema", "s", "", "Path to schema file")
	registerCmd.MarkFlagRequired("adapter")
	registerCmd.MarkFlagRequired("permissions")
	registerCmd.MarkFlagRequired("schema")
}
