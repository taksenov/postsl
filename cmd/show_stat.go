/*
Copyright © 2025 taksenov@gmail.com
*/

// Package cmd -- comand line interface app.
package cmd

import (
	"devops/app/statistics"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ShowStat -- showStatCmd pseudo constructor.
func ShowStat() *cobra.Command {
	showStatCmd := &cobra.Command{
		Use:   "show_stat",
		Short: "Отобразть статистику",
		Long: `Сгенерировать документацию отображающую статистику по пайплайнам.

Пример использования:
` + string("\033[92m") + `
  go run main.go show_stat --src '/путь/источника/' [> file_name.md]'
  go run main.go show_stat -s '/путь/источника/' [> file_name.md]'
` + string("\033[0m"),

		Run: func(_ *cobra.Command, _ []string) {
			statistics.ShowStatRunner()
		},
	}

	// local flags
	showStatCmd.PersistentFlags().StringP("src", "s", "", "Source dir")
	viper.BindPFlag("src.showStat", showStatCmd.PersistentFlags().Lookup("src"))

	return showStatCmd
}
