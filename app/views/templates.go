package views

import (
	"html/template"
	"io"
	"time"
)

var templateFuncMap = map[string]interface{}{
	"Format": func(t time.Time) string { return t.Format("2006-01-02") },
}

func GetTemplates(files ...string) *template.Template {
	return template.Must(template.New("").Funcs(templateFuncMap).ParseFiles(files...))
}

func RenderTemplate(t *template.Template, wr io.Writer, name string, data interface{}) error {
	return t.ExecuteTemplate(wr, name, data)
}
