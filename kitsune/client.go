package kitsune

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func Put(url string, body io.Reader) (re *http.Response, err error) {
	r, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return
	}
	r.Header.Set("Content-Type", "application/octet-stream")

	return http.DefaultClient.Do(r)
}

func PostJSON(url string, req any) (re *http.Response, err error) {
	r, err := NewJSONRequestWithContext(context.Background(), http.MethodPost, url, req)
	if err != nil {
		return
	}
	r.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(r)
}

func Delete(url string) (re *http.Response, err error) {
	r, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return
	}

	return http.DefaultClient.Do(r)
}

func NewJSONRequestWithContext(ctx context.Context, method string, url string, req any) (r *http.Request, err error) {
	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(req); err != nil {
		return
	}

	return http.NewRequestWithContext(ctx, method, url, &buf)
}
