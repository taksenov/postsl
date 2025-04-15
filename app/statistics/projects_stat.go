/*
Copyright © 2025 taksenov@gmail.com
*/

// Package statistics -- pipeline statistics.
package statistics

import (
	"devops/app/filesystem"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

type projectInfo struct {
	keyName       string
	projectsCount int32
	projectsList  []string
}

type ByName []projectInfo

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].keyName < a[j].keyName }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func showProjectsStat(src string) error {
	projects := make(map[string]projectInfo)

	manifests, err := filesystem.GetManifestNames(src)
	if err != nil {
		return err
	}

	for _, manifest := range manifests {
		viper.SetConfigName(src + manifest)
		viper.SetConfigFile(src + manifest)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		if viper.IsSet("type") {
			pipeType := viper.GetString("type")

			if _, ok := projects[pipeType]; !ok {
				p := projectInfo{
					keyName:       pipeType,
					projectsCount: 1,
					projectsList:  []string{strings.TrimSuffix(manifest, ".yaml")},
				}
				projects[pipeType] = p
			} else {
				p := projects[pipeType]
				p.projectsCount++
				p.projectsList = append(p.projectsList, strings.TrimSuffix(manifest, ".yaml"))
				projects[pipeType] = p
			}
		}
	}

	sortedProjects := make([]projectInfo, 0)
	for _, val := range projects {
		sortedProjects = append(sortedProjects, val)
	}
	sort.Stable(ByName(sortedProjects))

	fmt.Printf("### Projects stat\n\n")
	for _, v := range sortedProjects {
		fmt.Printf("Project type: **%s**\n\n", v.keyName)
		fmt.Printf("count: %d\n\n", v.projectsCount)
		fmt.Println("```")
		for _, el := range v.projectsList {
			fmt.Println(el)
		}
		fmt.Println("```")
		fmt.Printf("---\n\n")
	}

	return nil
}
