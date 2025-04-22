/*
Copyright © 2025 taksenov@gmail.com
*/

// Package cmd -- comand line interface app.
package cmd

import (
	"context"
	"devops/app/runner"

	"github.com/spf13/cobra"
)

// SshRun -- sshRunCmd pseudo constructor.
func SshRun(ctx context.Context) *cobra.Command {
	sshRunCmd := &cobra.Command{
		Use:   "ssh-run",
		Short: "Запустить команду на серверах",
		Long: `"Запустить команду на серверах.

Пример использования:
` + string("\033[92m") + `
  go run main.go ssh-run
` + string("\033[0m"),

		Run: func(_ *cobra.Command, _ []string) {
			runner.SshRunner(ctx)
		},
	}

	return sshRunCmd
}
