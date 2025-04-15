/*
Copyright © 2025 taksenov@gmail.com
*/

// Package cmd -- comand line interface app.
package cmd

import (
	"devops/app/generator"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GenPipe -- genPipeCmd pseudo constructor.
func GenPipe() *cobra.Command {
	genPipeCmd := &cobra.Command{
		Use:   "genPipe",
		Short: "Сгенерировать пайплайн на основе манифеста",
		Long: `Сгенерировать пайплайн на основе указанного манифеста.

Пример использования:
` + string("\033[92m") + `
  go run main.go genPipe --src '/путь/источника/' --dst '/путь/назначения/'
  go run main.go genPipe -s '/путь/источника/' -d '/путь/назначения/'
` + string("\033[0m"),

		Run: func(_ *cobra.Command, _ []string) {
			generator.GenPipeRunner()
		},
	}

	// local flags
	genPipeCmd.PersistentFlags().StringP("src", "s", "", "Source dir")
	viper.BindPFlag("src.genPipe", genPipeCmd.PersistentFlags().Lookup("src"))
	genPipeCmd.PersistentFlags().StringP("dst", "d", "", "Destination dir")
	viper.BindPFlag("dst.genPipe", genPipeCmd.PersistentFlags().Lookup("dst"))

	return genPipeCmd
}
