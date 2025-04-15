/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineSearch707 all fields for PipelineSearch707 pipeline.
type PipelineSearch707 struct {
	MainCommentGolang         string
	Include                   string
	StagesComment             string
	Stages                    string
	WorkflowComment           string
	Workflow                  string
	InterruptibleTemlate      string
	TemplatesComment          string
	GoImageTemplate           string
	DockerImageTemplate       string
	GolangciLintImageTemplate string
	BuildTemplate             string
	DockerizeTemplate         string
	CheckCodeQualityComment   string
	Test                      string
	Lint                      string
	Pages                     string
	StagingComment            string
	BuildStaging              string
	DockerizeStaging          string
	ProductionComment         string
	BuildProduction           string
	DockerizeProduction       string
}

// Generate for PipelineSearch707.
func (t *PipelineSearch707) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineGolang constructor.
func NewPipelineSearch707(data map[string]string) Pipeline {
	return &PipelineSearch707{
		MainCommentGolang:         data["MainCommentGolang"],
		Include:                   data["Include"],
		StagesComment:             data["StagesComment"],
		Stages:                    data["Stages"],
		WorkflowComment:           data["WorkflowComment"],
		Workflow:                  data["Workflow"],
		InterruptibleTemlate:      data["InterruptibleTemlate"],
		TemplatesComment:          data["TemplatesComment"],
		GoImageTemplate:           data["GoImageTemplate"],
		DockerImageTemplate:       data["DockerImageTemplate"],
		GolangciLintImageTemplate: data["GolangciLintImageTemplate"],
		DockerizeTemplate:         data["DockerizeTemplate"],
		BuildTemplate:             data["BuildTemplate"],
		CheckCodeQualityComment:   data["CheckCodeQualityComment"],
		Test:                      data["Test"],
		Lint:                      data["Lint"],
		Pages:                     data["Pages"],
		StagingComment:            data["StagingComment"],
		BuildStaging:              data["BuildStaging"],
		DockerizeStaging:          data["DockerizeStaging"],
		ProductionComment:         data["ProductionComment"],
		BuildProduction:           data["BuildProduction"],
		DockerizeProduction:       data["DockerizeProduction"],
	}
}
