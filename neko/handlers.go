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

func Ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Info("Pong!")
		if _, err := fmt.Fprintln(w, "Pong!"); err != nil {
			ls.Warning(err)
		}
	})
}

func AssertMethod(method string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			ls.Warning("Method not allowed.", r.Method, method)
			MethodNotAllowed(w)
			return
		}
		h.ServeHTTP(w, r)
	})
}
