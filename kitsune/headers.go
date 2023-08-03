package kitsune

import (
	"net/http"
	"strings"
)

func GetAuthorization(h http.Header) (token string, ok bool) {
	ss := strings.Split(h.Get("Authorization"), " ")
	if len(ss) != 2 || ss[0] != "Bearer" {
		return
	}
	return ss[1], true
}

func SetAuthorization(h http.Header, token string) {
	h.Set("Authorization", "Bearer "+token)
}
