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
	if resp == nil {
		ls.Warning("Empty response.")
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		ls.Error("Marshal response error.", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		ls.Warning("Write output error.", err)
	}
}

func Reject(w http.ResponseWriter, rej any, code int) {
	if rej == nil {
		http.Error(w, "", code)
		return
	}

	data, err := json.Marshal(rej)
	if err != nil {
		ls.Error("Marshal rejection error.", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(data); err != nil {
		ls.Warning("Write output error.", err)
	}
}
