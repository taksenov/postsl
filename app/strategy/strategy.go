/*
Copyright © 2025 taksenov@gmail.com
*/

// Package strategy -- design pattern strategy.
package strategy

// Pipeline interface.
type Pipeline interface {
	Generate(dst string, tmpl string) error
}

// Comment struct for YAML manifest.
type Comment struct {
	File string
	Key  string
}

// Comments all comments fields.
type Comments struct {
	MainCommentSSR          Comment
	MainCommentPHP          Comment
	MainCommentVendor       Comment
	MainCommentNonSSR       Comment
	MainCommentGolang       Comment
	VariablesComment        Comment
	StagesComment           Comment
	WorkflowComment         Comment
	TemplatesComment        Comment
	PrepareComment          Comment
	WarnCacheComment        Comment
	CheckCodeQualityComment Comment
	StagingComment          Comment
	ProductionComment       Comment
	FeaturesComment         Comment
	BuildComment            Comment
	DockerizeComment        Comment
	SentryComment           Comment
	CacheComment            Comment
	GenerateComment         Comment
	FormattingComment       Comment
	FinalizeComment         Comment
	PublishComment          Comment
}
