package controllers

import (
	"fmt"
	"net/http"

	"github.com/itang/gaemvc"
)

type WakeupAction struct {
	gaemvc.GaeAction
}

func (this *WakeupAction) Apply() {
	resp, err := this.HttpClient().Get("http://apps.itang.me")
	if err != nil {
		this.SendError(err.Error(), http.StatusInternalServerError)
		return
	}
	this.Send(fmt.Sprintf("HTTP GET returned status %v", resp.Status))
}
