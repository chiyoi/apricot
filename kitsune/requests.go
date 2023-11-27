package kitsune

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/chiyoi/iter/res"
)

func JSONReader(a any) (r io.Reader, err error) {
	var buf bytes.Buffer
	return &buf, json.NewEncoder(&buf).Encode(a)
}

func GetJSON(ctx context.Context, u string, auth func(r *http.Request) (*http.Request, error), a any) (err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	r, err = res.Then(r, err, auth)
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	defer re.Body.Close()
	return ParseResponse(re, a)
}

func GetStream(ctx context.Context, u string, auth func(r *http.Request) (*http.Request, error)) (body io.ReadCloser, err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	r, err = res.Then(r, err, auth)
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	return re.Body, ParseResponse(re, nil)
}

func PostJSON(ctx context.Context, u string, auth func(r *http.Request) (*http.Request, error), resp, req any) (err error) {
	body, err := JSONReader(req)
	r, err := res.Then(body, err, runnerNewRequestWithContext(ctx, http.MethodPost, u))
	r, err = res.Then(r, err, auth)
	r, err = res.Then(r, err, runnerSetHeader("Content-Type", "application/json"))
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	defer re.Body.Close()

	return ParseResponse(re, resp)
}

func PostStream(ctx context.Context, u string, auth func(r *http.Request) (*http.Request, error), body io.Reader, resp any) (err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
	r, err = res.Then(r, err, auth)
	r, err = res.Then(r, err, runnerSetHeader("Content-Type", "application/octet-stream"))
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	defer re.Body.Close()
	return ParseResponse(re, resp)
}

func Delete(ctx context.Context, u string, auth func(r *http.Request) (*http.Request, error)) (err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodDelete, u, nil)
	r, err = res.Then(r, err, auth)
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	defer re.Body.Close()
	return ParseResponse(re, nil)
}

func runnerNewRequestWithContext(ctx context.Context, method string, u string) func(body io.Reader) (*http.Request, error) {
	return func(body io.Reader) (*http.Request, error) {
		return http.NewRequestWithContext(ctx, http.MethodGet, u, body)
	}
}

func runnerSetHeader(key, val string) func(r *http.Request) (*http.Request, error) {
	return func(r *http.Request) (*http.Request, error) {
		r.Header.Set(key, val)
		return r, nil
	}
}
