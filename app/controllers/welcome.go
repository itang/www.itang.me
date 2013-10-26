package controllers

import (
	"net/http"

	"app/views"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	views.New(w).RenderTemplate("welcome.html", nil)
}
