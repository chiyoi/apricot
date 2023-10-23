package neko

import (
	"fmt"
	"net/http"
)

func TeapotHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Warning("I'm a teapot~")
		Teapot(w)
	})
}

func InternalServerErrorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Error("Internal server error.")
		InternalServerError(w)
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

func PingHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Info("Pong!")
		if _, err := fmt.Fprintln(w, "Pong!"); err != nil {
			ls.Warning(err)
		}
	})
}

func MethodAsserted(h http.Handler, method string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !GuardMethod(w, r, method) {
			return
		}
		h.ServeHTTP(w, r)
	})
}
