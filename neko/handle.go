package neko

import (
	"fmt"
	"net/http"
)

func BadRequest(w http.ResponseWriter, message string) {
	e(w, "400 bad request", http.StatusBadRequest, message)
}

func Unauthorized(w http.ResponseWriter, message string) {
	e(w, "401 unauthorized", http.StatusUnauthorized, message)
}

func Forbidden(w http.ResponseWriter, message string) {
	e(w, "403 forbidden", http.StatusForbidden, message)
}

func MethodNotAllowed(w http.ResponseWriter, message string) {
	e(w, "405 method not allowed", http.StatusMethodNotAllowed, message)
}

func Teapot(w http.ResponseWriter, message string) {
	e(w, "418 I'm a teapot", http.StatusTeapot, message)
}

func InternalServerError(w http.ResponseWriter, message string) {
	e(w, "500 internal server error", http.StatusInternalServerError, message)
}

func ServiceUnavailable(w http.ResponseWriter, message string) {
	e(w, "503 service unavailable", http.StatusServiceUnavailable, message)
}

func TemporaryRedirect(w http.ResponseWriter, r *http.Request, u string) {
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func PermanentRedirect(w http.ResponseWriter, r *http.Request, u string) {
	http.Redirect(w, r, u, http.StatusPermanentRedirect)
}

func e(w http.ResponseWriter, status string, code int, message string) {
	he := func(s string) { http.Error(w, s, code) }
	if message != "" {
		he(fmt.Sprintf("%s (%s)", status, message))
	} else {
		he(status)
	}
}
