/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineAcquiringfaker2023 all fields for PipelineAcquiringfaker2023 pipeline.
type PipelineAcquiringfaker2023 struct {
	MainCommentPHP      string
	Include             string
	StagesComment       string
	Stages              string
	WorkflowComment     string
	Workflow            string
	TemplatesComment    string
	DockerImageTemplate string
	RunnerTagsTemplate  string
	DockerizeTemplate   string
	BuildStaging        string
	StagingComment      string
	DockerizeStaging    string
}

// Generate for PipelineAcquiringfaker2023.
func (t *PipelineAcquiringfaker2023) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineAcquiringfaker2023 constructor.
func NewPipelineAcquiringfaker2023(data map[string]string) Pipeline {
	return &PipelineAcquiringfaker2023{
		MainCommentPHP:      data["MainCommentPHP"],
		Include:             data["Include"],
		StagesComment:       data["StagesComment"],
		Stages:              data["Stages"],
		WorkflowComment:     data["WorkflowComment"],
		Workflow:            data["Workflow"],
		TemplatesComment:    data["TemplatesComment"],
		DockerImageTemplate: data["DockerImageTemplate"],
		RunnerTagsTemplate:  data["RunnerTagsTemplate"],
		DockerizeTemplate:   data["DockerizeTemplate"],
		BuildStaging:        data["BuildStaging"],
		StagingComment:      data["StagingComment"],
		DockerizeStaging:    data["DockerizeStaging"],
	}
}
