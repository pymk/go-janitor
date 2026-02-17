package main

import (
	"github.com/alecthomas/kong"
	"github.com/pymk/go-janitor/cli"
)

func main() {
	var c cli.CLI

	// Parse CLI arguments with kong
	ctx := kong.Parse(&c,
		kong.Name("janitor"),
		kong.Description("String clean-up CLI tool for transforming stdin."),
		kong.UsageOnError(),
	)

	// Execute the transformation pipeline
	err := c.Run()

	// Handle errors
	ctx.FatalIfErrorf(err)
}
