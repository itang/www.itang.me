package filters

import (
	"fmt"
	"log"
	"net/http"

	"appengine"
)

func CatchErr(handler http.HandlerFunc) http.HandlerFunc {
	log.Println("Error filter")
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("500 Error:%v", err)
				log.Println(msg)

				appengine.NewContext(r).Errorf(msg)

				http.Error(w, msg, http.StatusInternalServerError)
			}
		}()

		handler(w, r)
	}
}
