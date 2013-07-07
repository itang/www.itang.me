package controllers

import (
  "appengine"
  "appengine/urlfetch"
  "fmt"
  "net/http"
)

func Wakeup(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  client := urlfetch.Client(c)
  resp, err := client.Get("http://apps.itang.me")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)
}
