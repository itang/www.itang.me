package controllers

import (
  "net/http"
)

import (
  v "app/views"
)

var (
  welcomeTmpl = v.GetTemplates("app/views/welcome.html")
)

func Welcome(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "text/html; charset=UTF-8")

  v.RenderTemplate(welcomeTmpl, w, "welcome.html", nil)
}
