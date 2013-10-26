package controllers

type WelcomeAction struct {
	Action
}

func (this *WelcomeAction) Apply() {
	this.Render().RenderTemplate("welcome.html", nil)
}
