package filters

import (
	"log"
	"net/http"

	"appengine"

	"app/utils"
)

func AppengineContext(handler http.HandlerFunc) http.HandlerFunc {
	log.Println("AppengineContext filter")
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Model.Context = appengine.NewContext(r)
		handler(w, r)
	}
}
