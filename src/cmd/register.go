/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/sestinj/basin-node/client"
	"github.com/sestinj/basin-node/util"
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
		// if !interactive {
		// 	fmt.Fprintln(os.Stderr, "This command can only be run in interactive mode. Use `basin attach` first.")
		// }
		adapter, err := cmd.Flags().GetString("adapter")
		permissions, err := cmd.Flags().GetString("permissions")
		schema, err := cmd.Flags().GetString("schema")
		url := args[0]
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse flags: %s", err.Error())
		}

		adata, err := util.UnmarshalFromFile[client.AdapterJson](adapter)
		pdata, err := util.UnmarshalFromFile[[]client.PermissionJson](permissions)
		sdata, err := util.UnmarshalFromFile[map[string]interface{}](schema)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read file: %s", err.Error())
		}

		cfg := client.NewConfiguration()
		apiClient := client.NewAPIClient(cfg)
		ctx := context.Background()

		registerRequest := client.NewRegisterRequest(url, *pdata, *adata, *sdata)

		ok, r, err := apiClient.DefaultApi.Register(ctx).RegisterRequest(*registerRequest).Execute()
		if err != nil || !ok {
			fmt.Fprintf(os.Stderr, "Failed to register resource: %s", err.Error())
		} else if r.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "Response returned error: %s", r.Status)
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
