/*
Copyright © 2025 taksenov@gmail.com
*/

// Package cmd -- comand line interface app.
package cmd

import (
	"context"
	"devops/app/customerrors"
	"devops/app/runner"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SshRun -- sshRunCmd pseudo constructor.
func SshRun(ctx context.Context) *cobra.Command {
	sshRunCmd := &cobra.Command{
		Use:   "ssh-run",
		Short: "Запустить команду на серверах",
		Long: `"Запустить команду на серверах.

Пример использования:
` + string("\033[92m") + `
  go run main.go ssh-run --conf 'имя файла конфигурации'
  go run main.go ssh-run -c 'имя файла конфигурации'
` + string("\033[0m"),

		Run: func(_ *cobra.Command, _ []string) {
			runner.SshRunner(ctx)
		},
	}

	// local flags
	sshRunCmd.PersistentFlags().StringP("conf", "c", "", "Configuration file")

	err := viper.BindPFlag("conf", sshRunCmd.PersistentFlags().Lookup("conf"))
	if err != nil {
		customerrors.HandleErr(fmt.Errorf("bind flag 'conf' error: %w", err), "SshRun")
	}

	return sshRunCmd
}
