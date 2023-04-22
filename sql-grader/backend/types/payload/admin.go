package payload

import "backend/types/embed"

type AdminLabImport struct {
	Code        *string            `yaml:"code" validate:"alphanum,lowercase"`
	Name        *string            `yaml:"name" validate:"required"`
	Description *string            `yaml:"description" validate:"required"`
	TemplateDb  *string            `yaml:"templateDb" validate:"required"`
	Generator   *string            `yaml:"generator" validate:"required"`
	Tasks       []*AdminTaskImport `yaml:"tasks" validate:"required"`
}

type AdminTaskImport struct {
	Code        *string         `json:"code" validate:"alphanum,lowercase"`
	Title       *string         `json:"title" validate:"required"`
	Description *string         `json:"description" validate:"required"`
	Tags        *embed.TaskTags `json:"tags" validate:"required"`
	Query       *string         `json:"query" validate:"required"`
	Hint        *string         `json:"hint" validate:"required"`
}
