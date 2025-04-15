/*
Copyright © 2025 taksenov@gmail.com
*/

// Package generator -- comand line interface app.
package generator

import (
	"devops/app/strategy"
)

const (
	ssr                = "SSR"
	nonSSR             = "nonSSR"
	golang             = "golang"
	search707          = "search707"
	hub1477            = "hub1477"
	apiv51737          = "apiv51737"
	vendor984          = "vendor984"
	acquiringfaker2023 = "acquiringfaker2023"
	steward1717        = "steward1717"
	sparta819          = "sparta819"
	ilium965           = "ilium965"
	openapi            = "openAPI"
)

var templates = map[string]string{
	ssr:                "__default__/template/ssr.tpl",
	nonSSR:             "__default__/template/nonssr.tpl",
	golang:             "__default__/template/golang.tpl",
	search707:          "__default__/template/search-707.tpl",
	apiv51737:          "__default__/template/api-v5-1737.tpl",
	vendor984:          "__default__/template/vendor-984.tpl",
	acquiringfaker2023: "__default__/template/acquiring-faker-2023.tpl",
	steward1717:        "__default__/template/steward-1717.tpl",
	sparta819:          "__default__/template/sparta-819.tpl",
	ilium965:           "__default__/template/ilium-965.tpl",
	openapi:            "__default__/template/openapi.tpl",
}

var ssrPipeScenarioJobs = strategy.PipelineSSR{
	Include:                             "__default__/common-include.yaml",
	Stages:                              "__default__/common-stages.yaml",
	Workflow:                            "__default__/common-workflow.yaml",
	Variables:                           "",
	NodeImageTemplate:                   "__default__/common-node-image-template.yaml",
	DockerImageTemplate:                 "__default__/common-docker-image-template.yaml",
	RunnerTagsTemplate:                  "__default__/common-runner-tags-template.yaml",
	CacheOnlyPullTemplate:               "__default__/common-cache-only-pull-template.yaml",
	SSRBuildTemplate:                    "__default__/ssr-build-template.yaml",
	DockerizeTemplate:                   "__default__/common-dockerize-template.yaml",
	SSRPublishAssetsTemplate:            "__default__/ssr-publish-assets-template.yaml",
	SSRFlushAssetsCDNTemplate:           "",
	SSRBuildVariablesTemplate:           "",
	SSRBuildVariablesStagingTemplate:    "",
	SSRBuildVariablesProductionTemplate: "",
	PipelineVariablesTemplate:           "__default__/common-pipeline-variables-template.yaml",
	RulesFeatureTemplate:                "",
	WarnCache:                           "__default__/common-warn-cache.yaml",
	Lint:                                "__default__/common-lint.yaml",
	Test:                                "__default__/common-test.yaml",
	TestDeploy:                          "__default__/common-test-deploy.yaml",
	TypeCheck:                           "",
	SSRBuildStaging:                     "__default__/ssr-build-staging.yaml",
	SSRPublishAssetsStaging:             "__default__/ssr-publish-assets-staging.yaml",
	FlushAssetsCDNStaging:               "",
	DockerizeStaging:                    "__default__/common-dockerize-staging.yaml",
	SSRBuildProduction:                  "__default__/ssr-build-production.yaml",
	SSRPublishAssetsProduction:          "__default__/ssr-publish-assets-production.yaml",
	FlushAssetsCDNProduction:            "",
	DockerizeProduction:                 "__default__/common-dockerize-production.yaml",
	BuildFeature:                        "",
	PublishFeaturesAssets:               "",
	DockerizeFeature:                    "",
}

var ssrPipeScenarioCommnents = strategy.Comments{
	MainCommentSSR: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_ssr",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	VariablesComment: strategy.Comment{
		File: "",
		Key:  "",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	WarnCacheComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "warn_cache",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	StagingComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "staging",
	},
	ProductionComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "production",
	},
	FeaturesComment: strategy.Comment{
		File: "",
		Key:  "",
	},
}

var nonSSRPipeScenarioJobs = strategy.PipelineNonSSR{
	Include:                                "__default__/common-include.yaml",
	Stages:                                 "__default__/common-stages.yaml",
	Workflow:                               "__default__/common-workflow.yaml",
	NodeImageTemplate:                      "__default__/common-node-image-template.yaml",
	DockerImageTemplate:                    "__default__/common-docker-image-template.yaml",
	RunnerTagsTemplate:                     "__default__/common-runner-tags-template.yaml",
	CacheOnlyPullTemplate:                  "__default__/common-cache-only-pull-template.yaml",
	NonSSRBuildTemplate:                    "__default__/nonssr-build-template.yaml",
	DockerizeTemplate:                      "__default__/common-dockerize-template.yaml",
	NonSSRPublishAssetsTemplate:            "__default__/nonssr-publish-assets-template.yaml",
	NonSSRFlushAssetsCDNTemplate:           "",
	NonSSRBuildVariablesTemplate:           "",
	NonSSRBuildVariablesStagingTemplate:    "",
	NonSSRBuildVariablesProductionTemplate: "",
	PipelineVariablesTemplate:              "__default__/common-pipeline-variables-template.yaml",
	WarnCache:                              "__default__/common-warn-cache.yaml",
	Lint:                                   "__default__/common-lint.yaml",
	Test:                                   "__default__/common-test.yaml",
	TestDeploy:                             "__default__/common-test-deploy.yaml",
	NonSSRBuildStaging:                     "__default__/nonssr-build-staging.yaml",
	NonSSRPublishAssetsStaging:             "__default__/nonssr-publish-assets-staging.yaml",
	FlushAssetsCDNStaging:                  "",
	DockerizeStaging:                       "__default__/common-dockerize-staging.yaml",
	NonSSRBuildProduction:                  "__default__/nonssr-build-production.yaml",
	NonSSRPublishAssetsProduction:          "__default__/nonssr-publish-assets-production.yaml",
	FlushAssetsCDNProduction:               "",
	PublishPackage:                         "",
	DockerizeProduction:                    "__default__/common-dockerize-production.yaml",
}

var nonSSRPipeScenarioCommnents = strategy.Comments{
	MainCommentNonSSR: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_nonssr",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	WarnCacheComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "warn_cache",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	StagingComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "staging",
	},
	ProductionComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "production",
	},
}

var golangPipeScenarioJobs = strategy.PipelineGolang{
	Include:                          "__default__/common-include.yaml",
	Stages:                           "__default__/golang-stages.yaml",
	Workflow:                         "__default__/common-workflow.yaml",
	Variables:                        "",
	RunnerSettingsTemplate:           "__default__/golang-runner-settings-template.yaml",
	GoImageTemplate:                  "__default__/golang-go-image-template.yaml",
	GolangciLintImageTemplate:        "__default__/golang-golangci-lint-image-template.yaml",
	SentryCliImageTemplate:           "__default__/common-sentry-cli-image-template.yaml",
	DockerImageTemplate:              "__default__/common-docker-image-template.yaml",
	CacheBeforeScriptTemplate:        "__default__/golang-cache-before-script-template.yaml",
	CacheConfigTemplate:              "__default__/golang-cache-config-template.yaml",
	UpdateCacheTemplate:              "__default__/golang-update-cache-template.yaml",
	CacheTemplate:                    "__default__/golang-cache-template.yaml",
	DockerizeTemplate:                "__default__/golang-dockerize-template.yaml",
	PreviousReleaseTemplate:          "__default__/golang-previous-release-template.yaml",
	SentryReleaseTemplate:            "",
	ReleaseVersionStagingTemplate:    "",
	ReleaseVersionProductionTemplate: "",
	Test:                             "__default__/golang-test.yaml",
	Lint:                             "__default__/golang-lint.yaml",
	Build:                            "__default__/golang-build.yaml",
	PreviousRelease:                  "__default__/golang-previous-release.yaml",
	PreviousReleaseTag:               "__default__/golang-previous-release-tag.yaml",
	DockerizeStaging:                 "__default__/golang-dockerize-staging.yaml",
	DockerizeProduction:              "__default__/golang-dockerize-production.yaml",
	SentryReleaseStaging:             "__default__/golang-sentry-release-staging.yaml",
	SentryReleaseProduction:          "__default__/golang-sentry-release-production.yaml",
	StoreCache:                       "__default__/golang-store-cache.yaml",
}

var golangPipeScenarioCommnents = strategy.Comments{
	MainCommentGolang: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_golang",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	VariablesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "variables",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	BuildComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "build",
	},
	DockerizeComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "dockerize",
	},
	SentryComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "sentry",
	},
	CacheComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "cache",
	},
}

var search707PipeScenarioJobs = strategy.PipelineSearch707{
	Include:                   "__default__/common-include.yaml",
	Stages:                    "__default__/search-707-stages.yaml",
	Workflow:                  "__default__/common-workflow.yaml",
	InterruptibleTemlate:      "__default__/search-707-interruptible-template.yaml",
	GoImageTemplate:           "__default__/golang-go-image-template.yaml",
	DockerImageTemplate:       "__default__/common-docker-image-template.yaml",
	GolangciLintImageTemplate: "__default__/golang-golangci-lint-image-template.yaml",
	DockerizeTemplate:         "__default__/search-707-dockerize-template.yaml",
	BuildTemplate:             "__default__/search-707-build-template.yaml",
	Test:                      "__default__/search-707-test.yaml",
	Lint:                      "__default__/search-707-lint.yaml",
	Pages:                     "__default__/search-707-pages.yaml",
	BuildStaging:              "__default__/search-707-build-staging.yaml",
	DockerizeStaging:          "__default__/search-707-dockerize-staging.yaml",
	BuildProduction:           "__default__/search-707-build-production.yaml",
	DockerizeProduction:       "__default__/search-707-dockerize-production.yaml",
}

var search707PipeScenarioCommnents = strategy.Comments{
	MainCommentGolang: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_golang",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	StagingComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "staging",
	},
	ProductionComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "production",
	},
}

var hub1477PipeScenarioJobs = strategy.PipelineHub1477{
	Stages:                    "__default__/hub-1477-stages.yaml",
	Workflow:                  "__default__/common-workflow.yaml",
	RunnerSettingsTemplate:    "__default__/hub-1477-runner-settings-template.yaml",
	GoImageTemplate:           "__default__/golang-go-image-template.yaml",
	GolangciLintImageTemplate: "__default__/golang-golangci-lint-image-template.yaml",
	DockerImageTemplate:       "__default__/common-docker-image-template.yaml",
	CacheVariablesTemplate:    "__default__/hub-1477-cache-variables-template.yaml",
	CacheConfigTemplate:       "__default__/hub-1477-cache-config-template.yaml",
	CacheTemplate:             "__default__/hub-1477-cache-template.yaml",
	DockerizeTemplate:         "__default__/hub-1477-dockerize-template.yaml",
	BuildTemplate:             "__default__/hub-1477-build-template.yaml",
	Test:                      "__default__/hub-1477-test.yaml",
	Lint:                      "__default__/hub-1477-lint.yaml",
	BuildStaging:              "__default__/hub-1477-build-staging.yaml",
	DockerizeStaging:          "__default__/hub-1477-dockerize-staging.yaml",
	BuildProduction:           "__default__/hub-1477-build-production.yaml",
	DockerizeProduction:       "__default__/hub-1477-dockerize-production.yaml",
}

var hub1477PipeScenarioCommnents = strategy.Comments{
	MainCommentGolang: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_golang",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	StagingComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "staging",
	},
	ProductionComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "production",
	},
}

var apiv51737PipeScenarioJobs = strategy.PipelineApiv51737{
	Stages:                    "",
	Workflow:                  "__default__/common-workflow.yaml",
	Variables:                 "",
	RunnerSettingsTemplate:    "",
	NodeImageTemplate:         "",
	GoImageTemplate:           "__default__/golang-go-image-template.yaml",
	DockerImageTemplate:       "__default__/common-docker-image-template.yaml",
	CacheBeforeScriptTemplate: "__default__/golang-cache-before-script-template.yaml",
	CacheConfigTemplate:       "__default__/golang-cache-config-template.yaml",
	CacheTemplate:             "__default__/golang-cache-template.yaml",
	DockerizeTemplate:         "",
	Test:                      "",
	Build:                     "",
	BuildHelp:                 "",
	DockerizeStaging:          "",
	DockerizeHelpStaging:      "",
	DockerizeProduction:       "",
	DockerizeHelpProduction:   "",
}

var apiv51737PipeScenarioCommnents = strategy.Comments{
	MainCommentGolang: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_golang",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	VariablesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "variables",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	BuildComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "build",
	},
	DockerizeComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "dockerize",
	},
	ProductionComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "production",
	},
}

var vendor984PipeScenarioJobs = strategy.PipelineVendor984{
	Include:                          "__default__/common-include.yaml",
	Stages:                           "__default__/common-stages.yaml",
	Workflow:                         "__default__/common-workflow.yaml",
	NodeImageTemplate:                "__default__/common-node-image-template.yaml",
	DockerImageTemplate:              "__default__/common-docker-image-template.yaml",
	RunnerTagsTemplate:               "__default__/common-runner-tags-template.yaml",
	CacheOnlyPullTemplate:            "",
	BuildTemplate:                    "",
	DockerizeTemplate:                "__default__/common-dockerize-template.yaml",
	PublishAssetsTemplate:            "",
	BuildVariablesTemplate:           "",
	BuildVariablesStagingTemplate:    "",
	BuildVariablesProductionTemplate: "",
	PipelineVariablesTemplate:        "__default__/common-pipeline-variables-template.yaml",
	WarnCache:                        "",
	Lint:                             "",
	Test:                             "",
	BuildStaging:                     "",
	PublishAssetsStaging:             "",
	DockerizeStaging:                 "__default__/common-dockerize-staging.yaml",
	BuildProduction:                  "",
	PublishAssetsProduction:          "",
	DockerizeProduction:              "__default__/common-dockerize-production.yaml",
}

var vendor984PipeScenarioCommnents = strategy.Comments{
	MainCommentVendor: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_vendor",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	WarnCacheComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "warn_cache",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	StagingComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "staging",
	},
	ProductionComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "production",
	},
}

var steward1717PipeScenarioJobs = strategy.PipelineSteward1717{
	Include:               "",
	Stages:                "",
	Workflow:              "__default__/common-workflow.yaml",
	DockerImageTemplate:   "__default__/common-docker-image-template.yaml",
	RunnerTagsTemplate:    "",
	DockerizeTemplate:     "",
	DependenciesTemplate:  "",
	BuildCheckCodeQuality: "",
	Test:                  "",
	BuildStaging:          "",
	DockerizeStaging:      "",
}

var steward1717PipeScenarioCommnents = strategy.Comments{
	MainCommentPHP: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_php",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	PrepareComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "prepare",
	},
	BuildComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "build",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	DockerizeComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "dockerize",
	},
}

var sparta819PipeScenarioJobs = strategy.PipelineSparta819{
	Stages:                  "__default__/golang-stages.yaml",
	Workflow:                "__default__/common-workflow.yaml",
	RunnerSettingsTemplate:  "",
	GoImageTemplate:         "",
	PhpImageTemplate:        "",
	SentryCliImageTemplate:  "__default__/common-sentry-cli-image-template.yaml",
	DockerImageTemplate:     "__default__/common-docker-image-template.yaml",
	DockerizeTemplate:       "",
	ComposerTemplate:        "",
	SentryReleaseTemplate:   "",
	OpenAPITemplate:         "",
	DepsPHP:                 "",
	DepsGo:                  "",
	LintPHP:                 "",
	TestPHP:                 "",
	TestGo:                  "",
	TestOpenAPI:             "",
	GenOpenAPI:              "",
	VendorPhar:              "",
	Build:                   "",
	DockerizeStaging:        "",
	DockerizeProduction:     "",
	SentryReleaseStaging:    "",
	SentryReleaseProduction: "",
	StoreCache:              "",
}

var sparta819PipeScenarioCommnents = strategy.Comments{
	MainCommentGolang: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_sparta",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	PrepareComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "prepare",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	BuildComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "build",
	},
	DockerizeComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "dockerize",
	},
	SentryComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "sentry",
	},
	CacheComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "cache",
	},
}

var acquiringfaker2023PipeScenarioJobs = strategy.PipelineAcquiringfaker2023{
	Include:             "__default__/common-include.yaml",
	Stages:              "",
	Workflow:            "__default__/common-workflow.yaml",
	DockerImageTemplate: "__default__/common-docker-image-template.yaml",
	RunnerTagsTemplate:  "__default__/golang-runner-settings-template.yaml",
	DockerizeTemplate:   "",
	BuildStaging:        "",
	StagingComment:      "",
	DockerizeStaging:    "",
}

var acquiringfaker2023PipeScenarioCommnents = strategy.Comments{
	MainCommentPHP: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_PHP",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	WarnCacheComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "warn_cache",
	},
	StagingComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "staging",
	},
	ProductionComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "production",
	},
}

var ilium965PipeScenarioJobs = strategy.PipelineIlium965{
	Include:                          "__default__/common-include.yaml",
	Stages:                           "__default__/golang-stages.yaml",
	Workflow:                         "__default__/common-workflow.yaml",
	Variables:                        "",
	RunnerSettingsTemplate:           "__default__/golang-runner-settings-template.yaml",
	GoImageTemplate:                  "__default__/golang-go-image-template.yaml",
	GolangciLintImageTemplate:        "__default__/golang-golangci-lint-image-template.yaml",
	SentryCliImageTemplate:           "__default__/common-sentry-cli-image-template.yaml",
	DockerImageTemplate:              "__default__/common-docker-image-template.yaml",
	CacheBeforeScriptTemplate:        "__default__/golang-cache-before-script-template.yaml",
	CacheConfigTemplate:              "__default__/golang-cache-config-template.yaml",
	UpdateCacheTemplate:              "__default__/golang-update-cache-template.yaml",
	CacheTemplate:                    "__default__/golang-cache-template.yaml",
	DockerizeTemplate:                "__default__/golang-dockerize-template.yaml",
	PreviousReleaseTemplate:          "__default__/golang-previous-release-template.yaml",
	SentryReleaseTemplate:            "",
	ReleaseVersionStagingTemplate:    "",
	ReleaseVersionProductionTemplate: "",
	Test:                             "__default__/golang-test.yaml",
	CheckDeferRowsClose:              "",
	Lint:                             "__default__/golang-lint.yaml",
	Build:                            "__default__/golang-build.yaml",
	PreviousRelease:                  "__default__/golang-previous-release.yaml",
	PreviousReleaseTag:               "__default__/golang-previous-release-tag.yaml",
	DockerizeStaging:                 "__default__/golang-dockerize-staging.yaml",
	DockerizeProduction:              "__default__/golang-dockerize-production.yaml",
	SentryReleaseStaging:             "__default__/golang-sentry-release-staging.yaml",
	SentryReleaseProduction:          "__default__/golang-sentry-release-production.yaml",
	StoreCache:                       "__default__/golang-store-cache.yaml",
}

var ilium965PipeScenarioCommnents = strategy.Comments{
	MainCommentGolang: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_golang",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	VariablesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "variables",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	BuildComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "build",
	},
	DockerizeComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "dockerize",
	},
	SentryComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "sentry",
	},
	CacheComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "cache",
	},
}

var openapiPipeScenarioJobs = strategy.PipelineOpenAPI{
	Stages:                    "__default__/open-api-stages.yaml",
	Workflow:                  "__default__/common-workflow.yaml",
	Variables:                 "",
	NodeImageTemplate:         "__default__/common-node-image-template.yaml",
	GeneratorImageTemplate:    "__default__/open-api-gen-image-template.yaml",
	GenTSImageTemplate:        "__default__/gen-ts-image-template.yaml",
	GoToolImageTemplate:       "__default__/open-api-go-tool-image-template.yaml",
	RunnerSettingsTemplate:    "__default__/open-api-runner-tags-template.yaml",
	TestSpecification:         "__default__/open-api-test-specification.yaml",
	GenerateGolangPackage:     "__default__/open-api-generate-golang-package.yaml",
	GenerateTypescriptPackage: "__default__/open-api-generate-typescript-package.yaml",
	GoFormatting:              "__default__/open-api-go-formatting.yaml",
	TypescriptFormatting:      "__default__/open-api-typescript-formatting.yaml",
	FinalizeChanges:           "__default__/open-api-finalize-changes.yaml",
	PublishTypescriptPackage:  "__default__/open-api-publish-typescript-package.yaml",
}

var openapiPipeScenarioCommnents = strategy.Comments{
	MainCommentGolang: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "main_openapi",
	},
	StagesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "stages",
	},
	WorkflowComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "workflow",
	},
	VariablesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "variables",
	},
	TemplatesComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "templates",
	},
	CheckCodeQualityComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "check_code_quality",
	},
	GenerateComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "generate",
	},
	FormattingComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "formatting",
	},
	FinalizeComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "finalize",
	},
	PublishComment: strategy.Comment{
		File: "__default__/comments.yaml",
		Key:  "publish",
	},
}
