/*
Copyright © 2025 taksenov@gmail.com
*/

// Package main -- comand line interface app.
package main

import (
	"context"
	"devops/cmd"
	"os"
)

func main() {
	ctx := context.Background()

	root := cmd.Root()
	root.AddCommand(cmd.SshRun(ctx))

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
