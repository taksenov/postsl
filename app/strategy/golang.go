/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineGolang all fields for Golang pipeline.
type PipelineGolang struct {
	MainCommentGolang                string
	Include                          string
	StagesComment                    string
	Stages                           string
	WorkflowComment                  string
	Workflow                         string
	VariablesComment                 string
	Variables                        string
	TemplatesComment                 string
	RunnerSettingsTemplate           string
	GoImageTemplate                  string
	GolangciLintImageTemplate        string
	SentryCliImageTemplate           string
	DockerImageTemplate              string
	CacheBeforeScriptTemplate        string
	CacheConfigTemplate              string
	UpdateCacheTemplate              string
	CacheTemplate                    string
	DockerizeTemplate                string
	PreviousReleaseTemplate          string
	SentryReleaseTemplate            string
	ReleaseVersionStagingTemplate    string
	ReleaseVersionProductionTemplate string
	CheckCodeQualityComment          string
	Test                             string
	Lint                             string
	BuildComment                     string
	Build                            string
	PreviousRelease                  string
	PreviousReleaseTag               string
	DockerizeComment                 string
	DockerizeStaging                 string
	DockerizeProduction              string
	SentryComment                    string
	SentryReleaseStaging             string
	SentryReleaseProduction          string
	CacheComment                     string
	StoreCache                       string
}

// Generate for PipelineGolang.
func (t *PipelineGolang) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineGolang constructor.
func NewPipelineGolang(data map[string]string) Pipeline {
	return &PipelineGolang{
		MainCommentGolang:                data["MainCommentGolang"],
		Include:                          data["Include"],
		StagesComment:                    data["StagesComment"],
		Stages:                           data["Stages"],
		WorkflowComment:                  data["WorkflowComment"],
		Workflow:                         data["Workflow"],
		VariablesComment:                 data["VariablesComment"],
		Variables:                        data["Variables"],
		TemplatesComment:                 data["TemplatesComment"],
		RunnerSettingsTemplate:           data["RunnerSettingsTemplate"],
		GoImageTemplate:                  data["GoImageTemplate"],
		GolangciLintImageTemplate:        data["GolangciLintImageTemplate"],
		SentryCliImageTemplate:           data["SentryCliImageTemplate"],
		DockerImageTemplate:              data["DockerImageTemplate"],
		CacheBeforeScriptTemplate:        data["CacheBeforeScriptTemplate"],
		CacheConfigTemplate:              data["CacheConfigTemplate"],
		UpdateCacheTemplate:              data["UpdateCacheTemplate"],
		CacheTemplate:                    data["CacheTemplate"],
		DockerizeTemplate:                data["DockerizeTemplate"],
		PreviousReleaseTemplate:          data["PreviousReleaseTemplate"],
		SentryReleaseTemplate:            data["SentryReleaseTemplate"],
		ReleaseVersionStagingTemplate:    data["ReleaseVersionStagingTemplate"],
		ReleaseVersionProductionTemplate: data["ReleaseVersionProductionTemplate"],
		CheckCodeQualityComment:          data["CheckCodeQualityComment"],
		Test:                             data["Test"],
		Lint:                             data["Lint"],
		BuildComment:                     data["BuildComment"],
		Build:                            data["Build"],
		PreviousRelease:                  data["PreviousRelease"],
		PreviousReleaseTag:               data["PreviousReleaseTag"],
		DockerizeComment:                 data["DockerizeComment"],
		DockerizeStaging:                 data["DockerizeStaging"],
		DockerizeProduction:              data["DockerizeProduction"],
		SentryComment:                    data["SentryComment"],
		SentryReleaseStaging:             data["SentryReleaseStaging"],
		SentryReleaseProduction:          data["SentryReleaseProduction"],
		CacheComment:                     data["CacheComment"],
		StoreCache:                       data["StoreCache"],
	}
}
