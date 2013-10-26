package controllers

import (
	"log"
	"net/http"

	"appengine"
	"appengine/urlfetch"
	"github.com/google/go-github/github"

	"app/views"
)

func Todo(w http.ResponseWriter, r *http.Request) {
	client := githubClient(r)

	issues, _, err := client.Issues.ListByRepo("itang", "todo.itang.me", nil)
	if err != nil {
		views.New(w).Render500(err)
	} else {
		log.Println("issues:", len(issues))
		views.New(w).RenderTemplate("todo.html", struct {
			Issues []github.Issue
		}{
			Issues: issues,
		})
	}
}

func githubClient(r *http.Request) *github.Client {
	httpClient := urlfetch.Client(appengine.NewContext(r))
	return github.NewClient(httpClient)
}
