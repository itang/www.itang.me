package views

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Render struct {
	writer  http.ResponseWriter
	request *http.Request
}

func New(w http.ResponseWriter, r *http.Request) *Render {
	return &Render{w, r}
}

func (this *Render) RenderTemplate(name string, data interface{}) error {
	this.SetHtmlContentType()
	return templates.ExecuteTemplate(this.writer, name, data)
}

func (this *Render) Render500(err error) {
	this.SetHtmlContentType().SetStatus(500)

	fmt.Fprintf(this.writer, "error:%v", err)
}

func (this *Render) SetHtmlContentType() *Render {
	this.SetContentType("text/html; charset=UTF-8")
	return this
}

func (this *Render) SetContentType(ct string) *Render {
	this.writer.Header().Set("Content-type", ct)
	return this
}

func (this *Render) SetStatus(status int) *Render {
	this.writer.WriteHeader(status)
	return this
}

// ---------------------------------------------------------------------------------------------------
var (
	templateFuncMap = map[string]interface{}{
		"Format": func(t time.Time) string { return t.Format("2006-01-02") },
	}

	templates = getTemplates("app/views/welcome.html",
		"app/views/todo.html")
)

func getTemplates(files ...string) *template.Template {
	return template.Must(template.New("").Funcs(templateFuncMap).ParseFiles(files...))
}
