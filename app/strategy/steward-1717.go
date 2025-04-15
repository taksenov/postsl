/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineSteward1717 all fields for steward-1717 pipeline.
type PipelineSteward1717 struct {
	MainCommentPHP          string
	Include                 string
	StagesComment           string
	Stages                  string
	WorkflowComment         string
	Workflow                string
	TemplatesComment        string
	DockerImageTemplate     string
	RunnerTagsTemplate      string
	DockerizeTemplate       string
	DependenciesTemplate    string
	PrepareComment          string
	BuildCheckCodeQuality   string
	CheckCodeQualityComment string
	Test                    string
	BuildComment            string
	BuildStaging            string
	DockerizeComment        string
	DockerizeStaging        string
}

// Generate for PipelineSteward1717.
func (t *PipelineSteward1717) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineSteward1717 constructor.
func NewPipelineSteward1717(data map[string]string) Pipeline {
	return &PipelineSteward1717{
		MainCommentPHP:          data["MainCommentPHP"],
		Include:                 data["Include"],
		StagesComment:           data["StagesComment"],
		Stages:                  data["Stages"],
		WorkflowComment:         data["WorkflowComment"],
		Workflow:                data["Workflow"],
		TemplatesComment:        data["TemplatesComment"],
		DockerImageTemplate:     data["DockerImageTemplate"],
		RunnerTagsTemplate:      data["RunnerTagsTemplate"],
		DockerizeTemplate:       data["DockerizeTemplate"],
		DependenciesTemplate:    data["DependenciesTemplate"],
		PrepareComment:          data["PrepareComment"],
		BuildCheckCodeQuality:   data["BuildCheckCodeQuality"],
		CheckCodeQualityComment: data["CheckCodeQualityComment"],
		Test:                    data["Test"],
		BuildComment:            data["BuildComment"],
		BuildStaging:            data["BuildStaging"],
		DockerizeComment:        data["DockerizeComment"],
		DockerizeStaging:        data["DockerizeStaging"],
	}
}
