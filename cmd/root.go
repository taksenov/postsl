/*
Copyright © 2025 taksenov@gmail.com
*/

// Package cmd -- comand line interface app.
package cmd

import (
	"github.com/spf13/cobra"
)

// Root -- rootCmd pseudo constructor.
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "postsl",
		Short: "postsl ssh runner",
		Long:  "postsl приложение для запуска команд на серверах по ssh",
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return rootCmd
}
