/*
Copyright © 2022 Basin authors@basin.dev
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcutil/base58"
	didUtil "github.com/sestinj/basin-node/did"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Basin node",
	Long: `Authenticate with Basin node
		- basin auth add <DID>
		- basin auth forget <DID>
	`,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new keystore file for the given DID.",
	Long:  "Add a new keystore file for the given DID.",
	Run: func(cmd *cobra.Command, args []string) {
		err := didUtil.AuthLogin()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to login: %s", err.Error())
		}
	},
}

var forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Delete the keystore for the given DID",
	Long:  "Delete the keystore for the given DID",
	Run: func(cmd *cobra.Command, args []string) {
		did := args[0]

		err := didUtil.DeleteKeystore(did)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete keystore: %s", err.Error())
		}
		fmt.Fprintln(os.Stdout, "Successfully deleted keystore file.")
	},
}

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract and print the info from your keystore file.",
	Long:  "Extract and print the info from your keystore file.",
	Run: func(cmd *cobra.Command, args []string) {
		did := args[0]
		pw := args[1]

		priv, err := didUtil.ReadKeystore(did, pw)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read keystore file: %s\n", err.Error())
		}

		privBase58 := base58.Encode(priv)

		fmt.Fprintf(os.Stdout, "Private Key: %s\nPublic Key: %s\nDID: %s\n", string(privBase58), priv.Public(), did)
	},
}

var genKeyCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate and store a new key pair",
	Long:  "Generate and store a new key pair",
	Run: func(cmd *cobra.Command, args []string) {
		pw := args[0]

		pub, priv, err := didUtil.NewPrivateKey(pw)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to generate new keypair: %s\n", err.Error())
		}

		fmt.Fprintf(os.Stdout, "%s %s", string(pub), string(priv))
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(addCmd)
	authCmd.AddCommand(forgetCmd)
	authCmd.AddCommand(extractCmd)
	authCmd.AddCommand(genKeyCmd)
}
