package controllers

import (
	"net/http"

	"app/views"
)

var (
	welcomeTmpl = views.GetTemplates("app/views/welcome.html")
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=UTF-8")

	views.RenderTemplate(welcomeTmpl, w, "welcome.html", nil)
}
