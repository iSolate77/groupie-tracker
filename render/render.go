package render

import (
	"context"
	"net/http"
	"text/template"
)

type TemplateReader struct {
	templates *template.Template
}

func NewTemplateReader(pattern string) (*TemplateReader, error) {
	tmpls, err := template.New("").ParseGlob("templates/*.html")
	if err != nil {
		return nil, err
	}
	return &TemplateReader{templates: tmpls}, nil
}

func (r *TemplateReader) Render(ctx context.Context, w http.ResponseWriter, name string, data interface{}) error {
	return r.templates.ExecuteTemplate(w, name, data)
}
