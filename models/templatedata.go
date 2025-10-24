package models

import "bandb/src/forms"

type TemplateData struct {
	CSRFToken string
	Data      map[string]interface{}
	Form      *forms.Form
}
