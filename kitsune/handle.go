package kitsune

import (
	"encoding/json"
	"net/http"
)

func BadRequest(w http.ResponseWriter, rej any) {
	Reject(w, rej, http.StatusBadRequest)
}

func Unauthorized(w http.ResponseWriter, rej any) {
	Reject(w, rej, http.StatusUnauthorized)
}

func Forbidden(w http.ResponseWriter, rej any) {
	Reject(w, rej, http.StatusForbidden)
}

func Teapot(w http.ResponseWriter, rej any) {
	Reject(w, rej, http.StatusTeapot)
}

func InternalServerError(w http.ResponseWriter, rej any) {
	Reject(w, rej, http.StatusInternalServerError)
}

func Respond(w http.ResponseWriter, resp any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		ls.Warning(err)
	}
}

func Reject(w http.ResponseWriter, rej any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(rej); err != nil {
		ls.Warning(err)
	}
}
