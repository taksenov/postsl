/*
Copyright © 2025 taksenov@gmail.com
*/

// Package cmd -- comand line interface app.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "0.1.0"
	golang  = "1.24"
)

// VersionRun -- show version.
func VersionRun() *cobra.Command {
	versionRunCmd := &cobra.Command{
		Use:   "version",
		Short: fmt.Sprintf("version: %s", version),
		Long:  fmt.Sprintf("version %s, golang: %s", version, golang),

		// Run: func(_ *cobra.Command, _ []string) {},
	}

	return versionRunCmd
}
