package kitsune

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/chiyoi/iter/res"
)

func JSONReader(a any) (r io.Reader, err error) {
	var buf bytes.Buffer
	return &buf, json.NewEncoder(&buf).Encode(a)
}

func GetJSON(ctx context.Context, resp any, hook HookRequest) func(u string) (struct{}, error) {
	return func(endpoint string) (none struct{}, err error) {
		r, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if hook != nil {
			r, err = res.Then(r, err, hook)
		}
		re, err := res.Then(r, err, http.DefaultClient.Do)
		if err != nil {
			return
		}
		defer re.Body.Close()
		err = ParseResponse(re, resp)
		return
	}
}

func GetStream(ctx context.Context, u string, auth AuthFunc) (body io.ReadCloser, err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if auth != nil {
		r, err = res.Then(r, err, auth)
	}
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	return re.Body, ParseResponse(re, nil)
}

func PostJSON(ctx context.Context, u string, auth AuthFunc, resp, req any) (err error) {
	body, err := JSONReader(req)
	r, err := res.Then(body, err, runnerNewRequestWithContext(ctx, http.MethodPost, u))
	if auth != nil {
		r, err = res.Then(r, err, auth)
	}
	r, err = res.Then(r, err, runnerSetHeaderLine("Content-Type", "application/json"))
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	defer re.Body.Close()

	return ParseResponse(re, resp)
}

func PostStream(ctx context.Context, endpoint string, query url.Values, header http.Header, body io.Reader, resp any) (err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	r, err = res.Then(r, err, runnerSetQuery(query))
	r, err = res.Then(r, err, runnerSetHeader(header))
	r, err = res.Then(r, err, runnerSetHeaderLine("Content-Type", "application/octet-stream"))
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	defer re.Body.Close()
	return ParseResponse(re, resp)
}

func Delete(ctx context.Context, endpoint string, query url.Values, header http.Header) (err error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	r, err = res.Then(r, err, runnerSetQuery(query))
	r, err = res.Then(r, err, runnerSetHeader(header))
	re, err := res.Then(r, err, http.DefaultClient.Do)
	if err != nil {
		return
	}
	defer re.Body.Close()
	return ParseResponse(re, nil)
}

func runnerNewRequestWithContext(ctx context.Context, method string, u string) func(body io.Reader) (*http.Request, error) {
	return func(body io.Reader) (*http.Request, error) {
		return http.NewRequestWithContext(ctx, method, u, body)
	}
}

func runnerSetQuery(values url.Values) func(r *http.Request) (*http.Request, error) {
	return func(r *http.Request) (*http.Request, error) {
		var err error
		for k := range values {
			r, err = res.Then(r, err, runnerSetQueryLine(k, values.Get(k)))
		}
		return r, err
	}
}

func runnerSetQueryLine(key, value string) func(r *http.Request) (*http.Request, error) {
	return func(r *http.Request) (*http.Request, error) {
		q := r.URL.Query()
		q.Set(key, value)
		r.URL.RawQuery = q.Encode()
		return r, nil
	}
}

func runnerSetHeader(header http.Header) func(r *http.Request) (*http.Request, error) {
	return func(r *http.Request) (*http.Request, error) {
		var err error
		for k := range header {
			r, err = res.Then(r, err, runnerSetQueryLine(k, header.Get(k)))
		}
		return r, err
	}
}

func runnerSetHeaderLine(key, value string) func(r *http.Request) (*http.Request, error) {
	return func(r *http.Request) (*http.Request, error) {
		r.Header.Set(key, value)
		return r, nil
	}
}

type HookRequest func(*http.Request) (*http.Request, error)
