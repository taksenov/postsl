/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

import (
	"devops/app/filesystem"
)

// PipelineOpenAPI all fields for PipelineOpenapi pipeline.
type PipelineOpenAPI struct {
	MainCommentGolang         string
	StagesComment             string
	Stages                    string
	WorkflowComment           string
	Workflow                  string
	VariablesComment          string
	Variables                 string
	TemplatesComment          string
	NodeImageTemplate         string
	GeneratorImageTemplate    string
	GenTSImageTemplate        string
	GoToolImageTemplate       string
	RunnerSettingsTemplate    string
	CheckCodeQualityComment   string
	TestSpecification         string
	GenerateComment           string
	GenerateGolangPackage     string
	GenerateTypescriptPackage string
	FormattingComment         string
	GoFormatting              string
	TypescriptFormatting      string
	FinalizeComment           string
	FinalizeChanges           string
	PublishComment            string
	PublishTypescriptPackage  string
}

// Generate for PipelineOpenAPI.
func (t *PipelineOpenAPI) Generate(dst string, tmpl string) error {
	return filesystem.GeneratePipe(t, dst, tmpl)
}

// NewPipelineOpenAPI constructor.
func NewPipelineOpenAPI(data map[string]string) Pipeline {
	return &PipelineOpenAPI{
		MainCommentGolang:         data["MainCommentGolang"],
		StagesComment:             data["StagesComment"],
		Stages:                    data["Stages"],
		WorkflowComment:           data["WorkflowComment"],
		Workflow:                  data["Workflow"],
		VariablesComment:          data["VariablesComment"],
		Variables:                 data["Variables"],
		TemplatesComment:          data["TemplatesComment"],
		NodeImageTemplate:         data["NodeImageTemplate"],
		GeneratorImageTemplate:    data["GeneratorImageTemplate"],
		GenTSImageTemplate:        data["GenTSImageTemplate"],
		GoToolImageTemplate:       data["GoToolImageTemplate"],
		RunnerSettingsTemplate:    data["RunnerSettingsTemplate"],
		CheckCodeQualityComment:   data["CheckCodeQualityComment"],
		TestSpecification:         data["TestSpecification"],
		GenerateComment:           data["GenerateComment"],
		GenerateGolangPackage:     data["GenerateGolangPackage"],
		GenerateTypescriptPackage: data["GenerateTypescriptPackage"],
		FormattingComment:         data["FormattingComment"],
		GoFormatting:              data["GoFormatting"],
		TypescriptFormatting:      data["TypescriptFormatting"],
		FinalizeComment:           data["FinalizeComment"],
		FinalizeChanges:           data["FinalizeChanges"],
		PublishComment:            data["PublishComment"],
		PublishTypescriptPackage:  data["PublishTypescriptPackage"],
	}
}
