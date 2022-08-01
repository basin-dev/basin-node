package main

import (
	"context"

	cmd "github.com/sestinj/basin-node/cmd"
	. "github.com/sestinj/basin-node/node"
)

func main() {
	ctx := context.Background()

	StartEverything(ctx)

	// Run the CLI
	cmd.Execute()
}
