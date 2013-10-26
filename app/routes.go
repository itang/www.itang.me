package app

import (
	"net/http"

	"github.com/bmizerany/pat"

	"app/controllers"
)

func init() {
	m := pat.New()

	//m.Get("/", f.Full(c.Welcome))
	m.Get("/", http.HandlerFunc(controllers.Welcome))
	m.Get("/wakeup", http.HandlerFunc(controllers.Wakeup))
	m.Get("/todo", http.HandlerFunc(controllers.Todo))

	http.Handle("/", m)
}
