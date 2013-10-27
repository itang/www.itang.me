package controllers

import (
	"github.com/itang/gaemvc"
)

type WelcomeAction struct {
	gaemvc.GaeAction
}

func (this *WelcomeAction) Apply() {
	this.Render().RenderTemplate("welcome.html", nil)
}
