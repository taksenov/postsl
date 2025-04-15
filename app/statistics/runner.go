/*
Copyright © 2025 taksenov@gmail.com
*/

// Package statistics -- pipeline statistics.
package statistics

import (
	"devops/app/customerrors"
	"fmt"

	"github.com/spf13/viper"
)

// ShowStatRunner -- раннер отображающий статистику.
func ShowStatRunner() {
	srcDir := viper.GetString("src.showStat")
	if srcDir == "" {
		customerrors.HandleErr(fmt.Errorf("source dir not set! Use '-h' for help"), "GenPipeRunner")
	}

	//nolint:lll
	fmt.Print(`
# Pipeline Templates

Проект универсального пайплайна gitlab CI/CD

[Поcтановка задач на доработку](https://gitlab.sima-land.ru/dev-dep/dev/devops/pipeline-templates/-/wikis/Create-issues-or-merge-requests)

---
`)
	fmt.Print(`
## Full statistcs information

`)

	err := showProjectsStat(srcDir)
	if err != nil {
		customerrors.HandleErr(fmt.Errorf("showProjectsStat: %w", err), "GenPipeRunner")
	}
}
