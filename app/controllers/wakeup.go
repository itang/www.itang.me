package controllers

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/urlfetch"
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
