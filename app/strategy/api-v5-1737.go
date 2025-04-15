/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineApiv51737 all fields for PipelineApiv51737 pipeline.
type PipelineApiv51737 struct {
	MainCommentGolang         string
	StagesComment             string
	Stages                    string
	WorkflowComment           string
	Workflow                  string
	VariablesComment          string
	Variables                 string
	TemplatesComment          string
	RunnerSettingsTemplate    string
	NodeImageTemplate         string
	GoImageTemplate           string
	DockerImageTemplate       string
	CacheBeforeScriptTemplate string
	CacheConfigTemplate       string
	CacheTemplate             string
	DockerizeTemplate         string
	CheckCodeQualityComment   string
	Test                      string
	BuildComment              string
	Build                     string
	BuildHelp                 string
	DockerizeComment          string
	DockerizeStaging          string
	DockerizeHelpStaging      string
	ProductionComment         string
	DockerizeProduction       string
	DockerizeHelpProduction   string
}

// Generate for PipelineApiv51737.
func (t *PipelineApiv51737) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineApiv51737 constructor.
func NewPipelineApiv51737(data map[string]string) Pipeline {
	return &PipelineApiv51737{
		MainCommentGolang:         data["MainCommentGolang"],
		StagesComment:             data["StagesComment"],
		Stages:                    data["Stages"],
		WorkflowComment:           data["WorkflowComment"],
		Workflow:                  data["Workflow"],
		VariablesComment:          data["VariablesComment"],
		Variables:                 data["Variables"],
		RunnerSettingsTemplate:    data["RunnerSettingsTemplate"],
		NodeImageTemplate:         data["NodeImageTemplate"],
		TemplatesComment:          data["TemplatesComment"],
		GoImageTemplate:           data["GoImageTemplate"],
		DockerImageTemplate:       data["DockerImageTemplate"],
		CacheBeforeScriptTemplate: data["CacheBeforeScriptTemplate"],
		CacheConfigTemplate:       data["CacheConfigTemplate"],
		CacheTemplate:             data["CacheTemplate"],
		DockerizeTemplate:         data["DockerizeTemplate"],
		CheckCodeQualityComment:   data["CheckCodeQualityComment"],
		Test:                      data["Test"],
		BuildComment:              data["BuildComment"],
		Build:                     data["Build"],
		BuildHelp:                 data["BuildHelp"],
		DockerizeComment:          data["DockerizeComment"],
		DockerizeStaging:          data["DockerizeStaging"],
		DockerizeHelpStaging:      data["DockerizeHelpStaging"],
		ProductionComment:         data["ProductionComment"],
		DockerizeProduction:       data["DockerizeProduction"],
		DockerizeHelpProduction:   data["DockerizeHelpProduction"],
	}
}
