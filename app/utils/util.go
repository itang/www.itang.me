package utils

import (
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type FilterFunc func(http.HandlerFunc) http.HandlerFunc

func Comp(filters ...FilterFunc) FilterFunc {
	if len(filters) == 0 {
		panic("filters number must >= 1")
	}

	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var ret = handler
			for i := len(filters) - 1; i >= 0; i-- {
				ret = filters[i](ret)
			}
			ret(w, r)
		}
	}
}
