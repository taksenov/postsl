/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineVendor984 all fields for PipelineVendor984 pipeline.
type PipelineVendor984 struct {
	MainCommentVendor                string
	Include                          string
	StagesComment                    string
	Stages                           string
	WorkflowComment                  string
	Workflow                         string
	TemplatesComment                 string
	NodeImageTemplate                string
	DockerImageTemplate              string
	RunnerTagsTemplate               string
	CacheOnlyPullTemplate            string
	BuildTemplate                    string
	DockerizeTemplate                string
	PublishAssetsTemplate            string
	BuildVariablesTemplate           string
	BuildVariablesStagingTemplate    string
	BuildVariablesProductionTemplate string
	PipelineVariablesTemplate        string
	WarnCacheComment                 string
	WarnCache                        string
	CheckCodeQualityComment          string
	Lint                             string
	Test                             string
	StagingComment                   string
	BuildStaging                     string
	PublishAssetsStaging             string
	DockerizeStaging                 string
	ProductionComment                string
	BuildProduction                  string
	PublishAssetsProduction          string
	DockerizeProduction              string
}

// Generate for PipelineVendor984.
func (t *PipelineVendor984) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineVendor984 constructor.
func NewPipelineVendor984(data map[string]string) Pipeline {
	return &PipelineVendor984{
		MainCommentVendor:                data["MainCommentVendor"],
		Include:                          data["Include"],
		StagesComment:                    data["StagesComment"],
		Stages:                           data["Stages"],
		WorkflowComment:                  data["WorkflowComment"],
		Workflow:                         data["Workflow"],
		TemplatesComment:                 data["TemplatesComment"],
		NodeImageTemplate:                data["NodeImageTemplate"],
		DockerImageTemplate:              data["DockerImageTemplate"],
		RunnerTagsTemplate:               data["RunnerTagsTemplate"],
		CacheOnlyPullTemplate:            data["CacheOnlyPullTemplate"],
		BuildTemplate:                    data["BuildTemplate"],
		DockerizeTemplate:                data["DockerizeTemplate"],
		PublishAssetsTemplate:            data["PublishAssetsTemplate"],
		BuildVariablesTemplate:           data["BuildVariablesTemplate"],
		BuildVariablesStagingTemplate:    data["BuildVariablesStagingTemplate"],
		BuildVariablesProductionTemplate: data["BuildVariablesProductionTemplate"],
		PipelineVariablesTemplate:        data["PipelineVariablesTemplate"],
		WarnCacheComment:                 data["WarnCacheComment"],
		WarnCache:                        data["WarnCache"],
		CheckCodeQualityComment:          data["CheckCodeQualityComment"],
		Lint:                             data["Lint"],
		Test:                             data["Test"],
		StagingComment:                   data["StagingComment"],
		BuildStaging:                     data["BuildStaging"],
		PublishAssetsStaging:             data["PublishAssetsStaging"],
		DockerizeStaging:                 data["DockerizeStaging"],
		ProductionComment:                data["ProductionComment"],
		BuildProduction:                  data["BuildProduction"],
		PublishAssetsProduction:          data["PublishAssetsProduction"],
		DockerizeProduction:              data["DockerizeProduction"],
	}
}
