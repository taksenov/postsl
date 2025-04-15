/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineSSR all fields for SSR pipeline.
type PipelineSSR struct {
	MainCommentSSR                      string
	Include                             string
	StagesComment                       string
	Stages                              string
	WorkflowComment                     string
	Workflow                            string
	VariablesComment                    string
	Variables                           string
	TemplatesComment                    string
	NodeImageTemplate                   string
	DockerImageTemplate                 string
	RunnerTagsTemplate                  string
	CacheOnlyPullTemplate               string
	SSRBuildTemplate                    string
	DockerizeTemplate                   string
	SSRPublishAssetsTemplate            string
	SSRFlushAssetsCDNTemplate           string
	SSRBuildVariablesTemplate           string
	SSRBuildVariablesStagingTemplate    string
	SSRBuildVariablesProductionTemplate string
	PipelineVariablesTemplate           string
	RulesFeatureTemplate                string
	WarnCacheComment                    string
	WarnCache                           string
	CheckCodeQualityComment             string
	Lint                                string
	Test                                string
	TestDeploy                          string
	TypeCheck                           string
	StagingComment                      string
	SSRBuildStaging                     string
	SSRPublishAssetsStaging             string
	FlushAssetsCDNStaging               string
	DockerizeStaging                    string
	ProductionComment                   string
	SSRBuildProduction                  string
	SSRPublishAssetsProduction          string
	FlushAssetsCDNProduction            string
	DockerizeProduction                 string
	FeaturesComment                     string
	BuildFeature                        string
	PublishFeaturesAssets               string
	DockerizeFeature                    string
}

// Generate for PipelineSSR.
func (t *PipelineSSR) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineSSR constructor.
func NewPipelineSSR(data map[string]string) Pipeline {
	return &PipelineSSR{
		MainCommentSSR:                      data["MainCommentSSR"],
		Include:                             data["Include"],
		StagesComment:                       data["StagesComment"],
		Stages:                              data["Stages"],
		WorkflowComment:                     data["WorkflowComment"],
		Workflow:                            data["Workflow"],
		VariablesComment:                    data["VariablesComment"],
		Variables:                           data["Variables"],
		TemplatesComment:                    data["TemplatesComment"],
		NodeImageTemplate:                   data["NodeImageTemplate"],
		DockerImageTemplate:                 data["DockerImageTemplate"],
		RunnerTagsTemplate:                  data["RunnerTagsTemplate"],
		CacheOnlyPullTemplate:               data["CacheOnlyPullTemplate"],
		SSRBuildTemplate:                    data["SSRBuildTemplate"],
		DockerizeTemplate:                   data["DockerizeTemplate"],
		SSRPublishAssetsTemplate:            data["SSRPublishAssetsTemplate"],
		SSRFlushAssetsCDNTemplate:           data["SSRFlushAssetsCDNTemplate"],
		SSRBuildVariablesTemplate:           data["SSRBuildVariablesTemplate"],
		SSRBuildVariablesStagingTemplate:    data["SSRBuildVariablesStagingTemplate"],
		SSRBuildVariablesProductionTemplate: data["SSRBuildVariablesProductionTemplate"],
		PipelineVariablesTemplate:           data["PipelineVariablesTemplate"],
		RulesFeatureTemplate:                data["RulesFeatureTemplate"],
		WarnCacheComment:                    data["WarnCacheComment"],
		WarnCache:                           data["WarnCache"],
		CheckCodeQualityComment:             data["CheckCodeQualityComment"],
		Lint:                                data["Lint"],
		Test:                                data["Test"],
		TestDeploy:                          data["TestDeploy"],
		TypeCheck:                           data["TypeCheck"],
		StagingComment:                      data["StagingComment"],
		SSRBuildStaging:                     data["SSRBuildStaging"],
		SSRPublishAssetsStaging:             data["SSRPublishAssetsStaging"],
		FlushAssetsCDNStaging:               data["FlushAssetsCDNStaging"],
		DockerizeStaging:                    data["DockerizeStaging"],
		ProductionComment:                   data["ProductionComment"],
		SSRBuildProduction:                  data["SSRBuildProduction"],
		SSRPublishAssetsProduction:          data["SSRPublishAssetsProduction"],
		FlushAssetsCDNProduction:            data["FlushAssetsCDNProduction"],
		DockerizeProduction:                 data["DockerizeProduction"],
		FeaturesComment:                     data["FeaturesComment"],
		BuildFeature:                        data["BuildFeature"],
		PublishFeaturesAssets:               data["PublishFeaturesAssets"],
		DockerizeFeature:                    data["DockerizeFeature"],
	}
}
