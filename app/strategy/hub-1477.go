/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineHub1477 all fields for PipelineHub1477 pipeline.
type PipelineHub1477 struct {
	MainCommentGolang         string
	StagesComment             string
	Stages                    string
	WorkflowComment           string
	Workflow                  string
	TemplatesComment          string
	RunnerSettingsTemplate    string
	GoImageTemplate           string
	GolangciLintImageTemplate string
	DockerImageTemplate       string
	CacheVariablesTemplate    string
	CacheConfigTemplate       string
	CacheTemplate             string
	DockerizeTemplate         string
	BuildTemplate             string
	CheckCodeQualityComment   string
	Test                      string
	Lint                      string
	StagingComment            string
	BuildStaging              string
	DockerizeStaging          string
	ProductionComment         string
	BuildProduction           string
	DockerizeProduction       string
}

// Generate for PipelineHub1477.
func (t *PipelineHub1477) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineHub1477 constructor.
func NewPipelineHub1477(data map[string]string) Pipeline {
	return &PipelineHub1477{
		MainCommentGolang:         data["MainCommentGolang"],
		StagesComment:             data["StagesComment"],
		Stages:                    data["Stages"],
		WorkflowComment:           data["WorkflowComment"],
		Workflow:                  data["Workflow"],
		TemplatesComment:          data["TemplatesComment"],
		RunnerSettingsTemplate:    data["RunnerSettingsTemplate"],
		GoImageTemplate:           data["GoImageTemplate"],
		GolangciLintImageTemplate: data["GolangciLintImageTemplate"],
		DockerImageTemplate:       data["DockerImageTemplate"],
		CacheVariablesTemplate:    data["CacheVariablesTemplate"],
		CacheConfigTemplate:       data["CacheConfigTemplate"],
		CacheTemplate:             data["CacheTemplate"],
		DockerizeTemplate:         data["DockerizeTemplate"],
		BuildTemplate:             data["BuildTemplate"],
		CheckCodeQualityComment:   data["CheckCodeQualityComment"],
		Test:                      data["Test"],
		Lint:                      data["Lint"],
		StagingComment:            data["StagingComment"],
		BuildStaging:              data["BuildStaging"],
		DockerizeStaging:          data["DockerizeStaging"],
		ProductionComment:         data["ProductionComment"],
		BuildProduction:           data["BuildProduction"],
		DockerizeProduction:       data["DockerizeProduction"],
	}
}
