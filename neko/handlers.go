package neko

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
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

func StatusHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ls.Info("Get status.")
		var buf bytes.Buffer
		fmt.Fprintln(&buf, "Nyan~")
		if v := os.Getenv("VERSION"); v != "" {
			fmt.Fprintf(&buf, "Version: %s\n", v)
		}
		if _, err := io.Copy(w, &buf); err != nil {
			ls.Warning("Write response error.", err)
		}
	})
}
