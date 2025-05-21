package models

type TemplateData struct {
	CSRFToken string
	Data      map[string]interface{}
}
