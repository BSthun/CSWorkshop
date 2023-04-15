package payload

type AdminLabImport struct {
	Code        *string            `json:"code" validate:"alphanum,uppercase"`
	Name        *string            `json:"name" validate:"required"`
	Description *string            `json:"description" validate:"required"`
	TemplateDb  *string            `json:"templateDb" validate:"required"`
	Tasks       []*AdminTaskImport `json:"tasks" validate:"required"`
}

type AdminTaskImport struct {
	Code        *string `json:"code" validate:"required"`
	Title       *string `json:"title" validate:"required"`
	Description *string `json:"description" validate:"required"`
	Query       *string `json:"query" validate:"required"`
	Hint        *string `json:"hint" validate:"required"`
}
