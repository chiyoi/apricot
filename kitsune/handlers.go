package kitsune

import (
	"net/http"
)

func TeapotHandler(rej any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Warning("I'm a teapot~", rej)
		Teapot(w, rej)
	})
}
