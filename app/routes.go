package app

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/itang/gaemvc"

	. "app/controllers"
)

func init() {
	m := pat.New()

	//m.Get("/", f.Full(c.Welcome))
	m.Get("/", &gaemvc.Handler{&WelcomeAction{}})
	m.Get("/wakeup", &gaemvc.Handler{&WakeupAction{}})
	m.Get("/todo", &gaemvc.Handler{&TodoAction{}})

	http.Handle("/", m)
}
