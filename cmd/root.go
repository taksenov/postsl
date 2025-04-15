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
		Use:   "devops",
		Short: "devops приложение для обслуживания pipeline-templates",
		Long: string("\033[92m") + `
▒███████████████████▒█████████████████████▒░
░▒▒ ▓░▒░▒░▒░▓░░░ ▒▒░▒ ░▓░▒ ▒ ░░ ▒░ ▒▒ ░ ░▓ ░
░░▒ ▒ ░ ▒░░ ▒░ ░ ▒░ ░  ▒ ░ ░  ░ ░  ░░   ░▒ ░
░ ░ ░ ░ ░ ░ ░    ░  ░  ░   ░    ░   ░    ░  
  ░ ░       ░░    ░      ░      ░   ░    ░  

      ██████ █████  ████  ███    ███       
        ██   ██    ██  ██ ████  ████       
        ██   ████  ██████ ██ ████ ██       
        ██   ██    ██  ██ ██  ██  ██       
        ██   █████ ██  ██ ██      ██       
                                           
                                           
█████  █████ ██   ██  █████  █████  ███████
██  ██ ██    ██   ██ ██   ██ ██  ██ ██     
██  ██ ████  ██   ██ ██   ██ █████  ███████
██  ██ ██     ██ ██  ██   ██ ██          ██
█████  █████   ███    █████  ██     ███████

  ░ ░       ░░    ░      ░      ░   ░    ░  
░ ░ ░ ░ ░ ░ ░    ░  ░  ░   ░    ░   ░    ░  
░░▒ ▒ ░ ▒░░ ▒░ ░ ▒░ ░  ▒ ░ ░  ░ ░  ░░   ░▒ ░
░▒▒ ▓░▒░▒░▒░▓░░░ ▒▒░▒ ░▓░▒ ▒ ░░ ▒░ ▒▒ ░ ░▓ ░
▒███████████████████▒█████████████████████▒░` + string("\033[0m"),
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return rootCmd
}
