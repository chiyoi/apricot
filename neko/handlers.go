package neko

import (
	"fmt"
	"net/http"
)

func TeapotHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Warning("I'm a teapot~", message)
		Teapot(w, message)
	})
}

func InternalServerErrorHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Error("Internal server error.", message)
		InternalServerError(w, message)
	})
}

func RedirectHandler(u string, code int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, u, code)
	})
}

func TemporaryRedirectHandler(u string) http.Handler {
	return RedirectHandler(u, http.StatusTemporaryRedirect)
}

func PermanentRedirectHandler(u string) http.Handler {
	return RedirectHandler(u, http.StatusPermanentRedirect)
}

func RedirectToSlashHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		TemporaryRedirect(w, r, r.URL.JoinPath("/").String())
	})
}

func WarmupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Info("Service warmup.")
		fmt.Fprintln(w, "Nyan~")
	})
}
