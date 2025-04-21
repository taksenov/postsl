/*
Copyright © 2025 taksenov@gmail.com
*/

// Package main -- comand line interface app.
package main

import (
	"devops/cmd"
	"os"
)

func main() {
	root := cmd.Root()
	root.AddCommand(cmd.SshRun())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
