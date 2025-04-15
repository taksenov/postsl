/*
Copyright © 2025 taksenov@gmail.com
*/

// Package generator -- pipeline generator.
package generator

import (
	"devops/app/customerrors"
	"devops/app/filesystem"
	"devops/app/strategy"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Генератор пайплайнов.
func GenPipeRunner() {
	srcDir := viper.GetString("src.genPipe")
	if srcDir == "" {
		customerrors.HandleErr(fmt.Errorf("source dir not set! Use '-h' for help"), "GenPipeRunner")
	}

	dstDir := viper.GetString("dst.genPipe")
	if dstDir == "" {
		customerrors.HandleErr(fmt.Errorf("destination dir not set! Use '-h' for help"), "GenPipeRunner")
	}

	manifests, err := filesystem.GetManifestNames(srcDir)
	if err != nil {
		customerrors.HandleErr(err, "GenPipeRunner")
	}

	for _, manifest := range manifests {
		viper.SetConfigName(srcDir + manifest)
		viper.SetConfigFile(srcDir + manifest)

		if err := viper.ReadInConfig(); err != nil {
			customerrors.HandleErr(err, "GenPipeRunner")
		}

		if viper.IsSet("type") {
			var pipe strategy.Pipeline
			var templateFile string

			pipeType := viper.Get("type")
			switch pipeType {
			case ssr:
				data := fillData(srcDir, ssrPipeScenarioJobs, ssrPipeScenarioCommnents)
				pipe = strategy.NewPipelineSSR(data)
				templateFile = templates[ssr]
			case nonSSR:
				data := fillData(srcDir, nonSSRPipeScenarioJobs, nonSSRPipeScenarioCommnents)
				pipe = strategy.NewPipelineNonSSR(data)
				templateFile = templates[nonSSR]
			case golang:
				data := fillData(srcDir, golangPipeScenarioJobs, golangPipeScenarioCommnents)
				pipe = strategy.NewPipelineGolang(data)
				templateFile = templates[golang]
			case search707:
				data := fillData(srcDir, search707PipeScenarioJobs, search707PipeScenarioCommnents)
				pipe = strategy.NewPipelineSearch707(data)
				templateFile = templates[search707]
			case hub1477:
				data := fillData(srcDir, hub1477PipeScenarioJobs, hub1477PipeScenarioCommnents)
				pipe = strategy.NewPipelineHub1477(data)
				templateFile = templates[hub1477]
			case apiv51737:
				data := fillData(srcDir, apiv51737PipeScenarioJobs, apiv51737PipeScenarioCommnents)
				pipe = strategy.NewPipelineApiv51737(data)
				templateFile = templates[apiv51737]
			case vendor984:
				data := fillData(srcDir, vendor984PipeScenarioJobs, vendor984PipeScenarioCommnents)
				pipe = strategy.NewPipelineVendor984(data)
				templateFile = templates[vendor984]
			case acquiringfaker2023:
				data := fillData(srcDir, acquiringfaker2023PipeScenarioJobs, acquiringfaker2023PipeScenarioCommnents)
				pipe = strategy.NewPipelineAcquiringfaker2023(data)
				templateFile = templates[acquiringfaker2023]
			case steward1717:
				data := fillData(srcDir, steward1717PipeScenarioJobs, steward1717PipeScenarioCommnents)
				pipe = strategy.NewPipelineSteward1717(data)
				templateFile = templates[steward1717]
			case sparta819:
				data := fillData(srcDir, sparta819PipeScenarioJobs, sparta819PipeScenarioCommnents)
				pipe = strategy.NewPipelineSparta819(data)
				templateFile = templates[sparta819]
			case ilium965:
				data := fillData(srcDir, ilium965PipeScenarioJobs, ilium965PipeScenarioCommnents)
				pipe = strategy.NewPipelineIlium965(data)
				templateFile = templates[ilium965]
			case openapi:
				data := fillData(srcDir, openapiPipeScenarioJobs, openapiPipeScenarioCommnents)
				pipe = strategy.NewPipelineOpenAPI(data)
				templateFile = templates[openapi]
			default:
				fmt.Println("Nothing to do.")
				return
			}

			processPipeline(dstDir+manifest, srcDir+templateFile, pipe)
			fmt.Println("Manifest '" + manifest + "' Processing status: OK")
		} else {
			customerrors.HandleErr(fmt.Errorf("manifest field 'type' is required in "+manifest), "GenPipeRunner")
		}
	}
}

func processPipeline(dst string, tmpl string, pipe strategy.Pipeline) {
	err := pipe.Generate(dst, tmpl)
	if err != nil {
		customerrors.HandleErr(err, "processPipeline")
	}
}

func fillData(src string, jobs interface{}, comments interface{}) map[string]string {
	res := make(map[string]string, 0)

	var custom interface{}
	nullish := make(map[string]string)
	if viper.IsSet("custom") {
		custom = viper.Get("custom")

		// Make Nullish
		for key, val := range custom.(map[string]interface{}) {
			if val == "" {
				nullish[key] = ""
			}
		}
	} else {
		customerrors.HandleErr(fmt.Errorf("manifest field 'custom' is required"), "fillData")
	}

	// JOBS
	values := reflect.ValueOf(jobs)
	typesOf := values.Type()
	for i := 0; i < values.NumField(); i++ {
		key := typesOf.Field(i).Name
		val := values.Field(i).String()

		// From custom YAML manifest
		customData, ok := custom.(map[string]interface{})[strings.ToLower(key)]
		if ok {
			res[key] = customData.(string)
			continue
		}

		if val == "" {
			continue
		}

		// From __default__
		data, err := filesystem.ReadFileByLines(src + val)
		if err != nil {
			customerrors.HandleErr(err, "fillData")
		}
		res[key] = data

		// Check for Nullish
		if checkExist(&nullish, strings.ToLower(key)) {
			res[key] = ""
		}
	}

	// COMMENTS
	values = reflect.ValueOf(comments)
	typesOf = values.Type()
	for i := 0; i < values.NumField(); i++ {
		key := typesOf.Field(i).Name

		val := values.Field(i).Interface()
		file := val.(strategy.Comment).File
		yamlKey := val.(strategy.Comment).Key
		if file == "" {
			continue
		}
		if yamlKey == "" {
			continue
		}

		viper.SetConfigName(src + file)
		viper.SetConfigFile(src + file)

		if err := viper.ReadInConfig(); err != nil {
			customerrors.HandleErr(err, "fillData")
		}

		if viper.IsSet(yamlKey) {
			data := viper.Get(yamlKey)
			res[key] = data.(string)

			// Check for Nullish
			if checkExist(&nullish, strings.ToLower(key)) {
				res[key] = ""
			}
		} else {
			customerrors.HandleErr(fmt.Errorf("comment manifest field '"+yamlKey+"' is required"), "fillData")
		}
	}

	return res
}

func checkExist(c *map[string]string, k string) bool {
	_, res := (*c)[k]
	return res
}
