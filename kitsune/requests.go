package kitsune

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chiyoi/iter/res"
)

func NewJSONRequestWithContext(ctx context.Context, method string, url string, req any) (r *http.Request, err error) {
	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(req); err != nil {
		return
	}

	return http.NewRequestWithContext(ctx, method, url, &buf)
}

func Get(ctx context.Context, u string) (*http.Response, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	return res.Then(r, err, http.DefaultClient.Do)
}

func GetJSON(ctx context.Context, u string, a any) (err error) {
	re, err := Get(ctx, u)
	if err != nil {
		return
	}
	defer re.Body.Close()
	return ParseResponse(re, a)
}

func PostStream(ctx context.Context, u string, body io.Reader, resp any) (err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	if err != nil {
		return
	}

	r.Header.Set("Content-Type", "application/octet-stream")
	re, err := http.DefaultClient.Do(r)
	if err != nil {
		return
	}
	defer re.Body.Close()

	return ParseResponse(re, resp)
}

func PostJSON(ctx context.Context, u string, resp, req any) (err error) {
	r, err := NewJSONRequestWithContext(ctx, http.MethodPost, u, req)
	if err != nil {
		return
	}

	r.Header.Set("Content-Type", "application/json")
	re, err := http.DefaultClient.Do(r)
	if err != nil {
		return
	}
	defer re.Body.Close()

	return ParseResponse(re, resp)
}

func Delete(ctx context.Context, u string) (response *http.Response, err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodDelete, u, nil)
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	defer re.Body.Close()
	return re, ParseResponse(re, nil)
}
