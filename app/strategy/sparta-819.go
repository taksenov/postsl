/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineSparta819 all fields for sparta-819 pipeline.
type PipelineSparta819 struct {
	MainCommentGolang       string
	StagesComment           string
	Stages                  string
	WorkflowComment         string
	Workflow                string
	VariablesComment        string
	TemplatesComment        string
	RunnerSettingsTemplate  string
	GoImageTemplate         string
	PhpImageTemplate        string
	SentryCliImageTemplate  string
	DockerImageTemplate     string
	ComposerTemplate        string
	DockerizeTemplate       string
	SentryReleaseTemplate   string
	OpenAPITemplate         string
	PrepareComment          string
	DepsPHP                 string
	DepsGo                  string
	CheckCodeQualityComment string
	LintPHP                 string
	TestPHP                 string
	TestGo                  string
	TestOpenAPI             string
	GenOpenAPI              string
	BuildComment            string
	VendorPhar              string
	Build                   string
	DockerizeComment        string
	DockerizeStaging        string
	DockerizeProduction     string
	SentryComment           string
	SentryReleaseStaging    string
	SentryReleaseProduction string
	CacheComment            string
	StoreCache              string
}

// Generate for PipelineSparta819.
func (t *PipelineSparta819) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineSparta819 constructor.
func NewPipelineSparta819(data map[string]string) Pipeline {
	return &PipelineSparta819{
		MainCommentGolang:       data["MainCommentGolang"],
		StagesComment:           data["StagesComment"],
		Stages:                  data["Stages"],
		WorkflowComment:         data["WorkflowComment"],
		Workflow:                data["Workflow"],
		VariablesComment:        data["VariablesComment"],
		TemplatesComment:        data["TemplatesComment"],
		RunnerSettingsTemplate:  data["RunnerSettingsTemplate"],
		GoImageTemplate:         data["GoImageTemplate"],
		PhpImageTemplate:        data["PhpImageTemplate"],
		SentryCliImageTemplate:  data["SentryCliImageTemplate"],
		DockerImageTemplate:     data["DockerImageTemplate"],
		ComposerTemplate:        data["ComposerTemplate"],
		DockerizeTemplate:       data["DockerizeTemplate"],
		SentryReleaseTemplate:   data["SentryReleaseTemplate"],
		OpenAPITemplate:         data["OpenAPITemplate"],
		CheckCodeQualityComment: data["CheckCodeQualityComment"],
		PrepareComment:          data["PrepareComment"],
		DepsPHP:                 data["DepsPHP"],
		DepsGo:                  data["DepsGo"],
		TestPHP:                 data["TestPHP"],
		LintPHP:                 data["LintPHP"],
		TestGo:                  data["TestGo"],
		TestOpenAPI:             data["TestOpenAPI"],
		GenOpenAPI:              data["GenOpenAPI"],
		BuildComment:            data["BuildComment"],
		VendorPhar:              data["VendorPhar"],
		Build:                   data["Build"],
		DockerizeComment:        data["DockerizeComment"],
		DockerizeStaging:        data["DockerizeStaging"],
		DockerizeProduction:     data["DockerizeProduction"],
		SentryComment:           data["SentryComment"],
		SentryReleaseStaging:    data["SentryReleaseStaging"],
		SentryReleaseProduction: data["SentryReleaseProduction"],
		CacheComment:            data["CacheComment"],
		StoreCache:              data["StoreCache"],
	}
}
