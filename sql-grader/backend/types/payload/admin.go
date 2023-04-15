package payload

import "backend/types/embed"

type AdminLabImport struct {
	Code        *string            `json:"code" validate:"alphanum,uppercase"`
	Name        *string            `json:"name" validate:"required"`
	Description *string            `json:"description" validate:"required"`
	TemplateDb  *string            `json:"templateDb" validate:"required"`
	Tasks       []*AdminTaskImport `json:"tasks" validate:"required"`
}

type AdminTaskImport struct {
	Code        *string         `json:"code" validate:"alphanum,uppercase"`
	Title       *string         `json:"title" validate:"required"`
	Description *string         `json:"description" validate:"required"`
	Tags        *embed.TaskTags `json:"tags" validate:"required"`
	Query       *string         `json:"query" validate:"required"`
	Hint        *string         `json:"hint" validate:"required"`
}
