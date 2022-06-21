package main

import (
 "fmt"
 "os"
 cobra "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:  "basin-node",
    Short: "basin-node - a CLI for interacting with your Basin Node.",
    Long: `basin-node is a CLI that lets you do everything possible with a Basin Node.
   
One can use basin-node to manage producer, consumer, & user functionality`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}

func main() {
    Execute()
}