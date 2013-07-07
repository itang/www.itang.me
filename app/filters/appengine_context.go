package filters

import (
  "appengine"
  "log"
  "net/http"
)

import u "app/utils"

func AppengineContext(handler http.HandlerFunc) http.HandlerFunc {
  log.Println("AppengineContext filter")
  return func(w http.ResponseWriter, r *http.Request) {
    u.Model.Context = appengine.NewContext(r)
    handler(w, r)
  }
}
