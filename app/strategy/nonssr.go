/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineNonSSR all fields for nonSSR pipeline.
type PipelineNonSSR struct {
	MainCommentNonSSR                      string
	Include                                string
	StagesComment                          string
	Stages                                 string
	WorkflowComment                        string
	Workflow                               string
	TemplatesComment                       string
	NodeImageTemplate                      string
	DockerImageTemplate                    string
	RunnerTagsTemplate                     string
	CacheOnlyPullTemplate                  string
	NonSSRBuildTemplate                    string
	DockerizeTemplate                      string
	NonSSRPublishAssetsTemplate            string
	NonSSRFlushAssetsCDNTemplate           string
	NonSSRBuildVariablesTemplate           string
	NonSSRBuildVariablesStagingTemplate    string
	NonSSRBuildVariablesProductionTemplate string
	PipelineVariablesTemplate              string
	WarnCacheComment                       string
	WarnCache                              string
	CheckCodeQualityComment                string
	Lint                                   string
	Test                                   string
	TestDeploy                             string
	StagingComment                         string
	NonSSRBuildStaging                     string
	NonSSRPublishAssetsStaging             string
	FlushAssetsCDNStaging                  string
	DockerizeStaging                       string
	ProductionComment                      string
	NonSSRBuildProduction                  string
	NonSSRPublishAssetsProduction          string
	FlushAssetsCDNProduction               string
	PublishPackage                         string
	DockerizeProduction                    string
}

// Generate for PipelineNonSSR.
func (t *PipelineNonSSR) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineNonSSR constructor.
func NewPipelineNonSSR(data map[string]string) Pipeline {
	return &PipelineNonSSR{
		MainCommentNonSSR:                      data["MainCommentNonSSR"],
		Include:                                data["Include"],
		StagesComment:                          data["StagesComment"],
		Stages:                                 data["Stages"],
		WorkflowComment:                        data["WorkflowComment"],
		Workflow:                               data["Workflow"],
		TemplatesComment:                       data["TemplatesComment"],
		NodeImageTemplate:                      data["NodeImageTemplate"],
		DockerImageTemplate:                    data["DockerImageTemplate"],
		RunnerTagsTemplate:                     data["RunnerTagsTemplate"],
		CacheOnlyPullTemplate:                  data["CacheOnlyPullTemplate"],
		NonSSRBuildTemplate:                    data["NonSSRBuildTemplate"],
		DockerizeTemplate:                      data["DockerizeTemplate"],
		NonSSRPublishAssetsTemplate:            data["NonSSRPublishAssetsTemplate"],
		NonSSRFlushAssetsCDNTemplate:           data["NonSSRFlushAssetsCDNTemplate"],
		NonSSRBuildVariablesTemplate:           data["NonSSRBuildVariablesTemplate"],
		NonSSRBuildVariablesStagingTemplate:    data["NonSSRBuildVariablesStagingTemplate"],
		NonSSRBuildVariablesProductionTemplate: data["NonSSRBuildVariablesProductionTemplate"],
		PipelineVariablesTemplate:              data["PipelineVariablesTemplate"],
		WarnCacheComment:                       data["WarnCacheComment"],
		WarnCache:                              data["WarnCache"],
		CheckCodeQualityComment:                data["CheckCodeQualityComment"],
		Lint:                                   data["Lint"],
		Test:                                   data["Test"],
		TestDeploy:                             data["TestDeploy"],
		StagingComment:                         data["StagingComment"],
		NonSSRBuildStaging:                     data["NonSSRBuildStaging"],
		NonSSRPublishAssetsStaging:             data["NonSSRPublishAssetsStaging"],
		FlushAssetsCDNStaging:                  data["FlushAssetsCDNStaging"],
		DockerizeStaging:                       data["DockerizeStaging"],
		ProductionComment:                      data["ProductionComment"],
		NonSSRBuildProduction:                  data["NonSSRBuildProduction"],
		NonSSRPublishAssetsProduction:          data["NonSSRPublishAssetsProduction"],
		FlushAssetsCDNProduction:               data["FlushAssetsCDNProduction"],
		PublishPackage:                         data["PublishPackage"],
		DockerizeProduction:                    data["DockerizeProduction"],
	}
}
