package app

import (
	"net/http"

	"github.com/bmizerany/pat"

	. "app/controllers"
)

func init() {
	m := pat.New()

	//m.Get("/", f.Full(c.Welcome))
	m.Get("/", &Handler{&WelcomeAction{}})
	m.Get("/wakeup", http.HandlerFunc(Wakeup))
	m.Get("/todo", &Handler{&TodoAction{}})

	http.Handle("/", m)
}
