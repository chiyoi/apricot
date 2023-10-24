package neko

import (
	"net/http"
)

func BadRequest(w http.ResponseWriter) {
	http.Error(w, "400 bad request", http.StatusBadRequest)
}

func Unauthorized(w http.ResponseWriter) {
	http.Error(w, "401 unauthorized", http.StatusUnauthorized)
}

func Forbidden(w http.ResponseWriter) {
	http.Error(w, "403 forbidden", http.StatusForbidden)
}

func MethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
}

func Teapot(w http.ResponseWriter) {
	http.Error(w, "418 I'm a teapot", http.StatusTeapot)
}

func InternalServerError(w http.ResponseWriter) {
	http.Error(w, "500 internal server error", http.StatusInternalServerError)
}

func ServiceUnavailable(w http.ResponseWriter) {
	http.Error(w, "503 service unavailable", http.StatusServiceUnavailable)
}

func TemporaryRedirect(w http.ResponseWriter, r *http.Request, u string) {
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func PermanentRedirect(w http.ResponseWriter, r *http.Request, u string) {
	http.Redirect(w, r, u, http.StatusPermanentRedirect)
}

func AssertMethod(w http.ResponseWriter, r *http.Request, method string) (ok bool) {
	if r.Method != method {
		ls.Warning("Method not allowed.", "r.Method:", r.Method, "method:", method)
		MethodNotAllowed(w)
		return
	}
	return true
}
