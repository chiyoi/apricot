package kitsune

import (
	"net/http"
	"strings"

	"github.com/chiyoi/iter/res"
)

func GetAuthorization(h http.Header) (token string, ok bool) {
	ss := strings.Split(h.Get("Authorization"), " ")
	if len(ss) != 2 || ss[0] != "Bearer" {
		return
	}
	return ss[1], true
}

func SetAuthorization(token string) res.Hook[*http.Request] {
	return func(r *http.Request) (*http.Request, error) {
		r.Header.Set("Authorization", "Bearer "+token)
		return r, nil
	}
}

func SetHeaderContentTypeJSON(r *http.Request) (*http.Request, error) {
	r.Header.Set("Content-Type", "application/json")
	return r, nil
}

func SetHeaderContentTypeStream(r *http.Request) (*http.Request, error) {
	r.Header.Set("Content-Type", "application/octet-stream")
	return r, nil
}
