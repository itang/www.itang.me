package app

import (
  "github.com/bmizerany/pat"
  "net/http"
)

import (
  c "app/controllers"
  //f "app/filters"
)

func init() {
  m := pat.New()

  //m.Get("/", f.Full(c.Welcome))
  m.Get("/", http.HandlerFunc(c.Welcome))
  m.Get("/wakeup", http.HandlerFunc(c.Wakeup))

  http.Handle("/", m)
}
