/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/sestinj/basin-node/log"

	"github.com/sestinj/basin-node/util"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
// basin gen types:ts <URL>
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate code from the schema of the given resource",
	Long: `Use the schema of a resource to generate code, including:
	- types
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if !interactive {
			fmt.Fprintln(os.Stderr, "This command can only be run in interactive mode. Use `basin attach` first.")
		}
		action := args[0]
		dataUrl := args[1]
		// out := args[2]

		ctx := context.Background()

		// Retrieve the schema
		url := util.GetMetadataUrl(dataUrl, util.Schema)
		resp, r, err := interactiveConfig.ApiClient.DefaultApi.Read(ctx).Url(url).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read resource: %s\n", err.Error())
			return
		} else if r.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "Failed to read resource: %s\n", r.Status)
			return
		}
		fmt.Fprintln(os.Stdout, resp)

		// Use it to generate the requested code to the path `out`
		switch action {
		case "types:ts":
			//
			// TODO: Should be able to request either raw binary or json. Can this just happen through MIME types?? This should also depend on the schema/data type. Not everything will be JSON.
		default:
			log.Error.Fatal("Unknown action: ", action)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
