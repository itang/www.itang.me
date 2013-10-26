package controllers

import (
	"log"

	"github.com/google/go-github/github"
)

type TodoAction struct {
	Action
}

func (this *TodoAction) Apply() {
	issues, _, err := this.GithubClient().Issues.ListByRepo("itang", "todo.itang.me", nil)
	if err != nil {
		this.Render().Render500(err)
	} else {
		log.Println("issues:", len(issues))
		this.Render().RenderTemplate("todo.html", struct {
			Issues []github.Issue
		}{
			Issues: issues,
		})
	}
}
