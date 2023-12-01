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

func Get(ctx context.Context, hook res.Hook[*http.Request]) func(u string) (body io.ReadCloser, err error) {
	return func(u string) (body io.ReadCloser, err error) {
		r, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
		if hook != nil {
			r, err = res.R(r, err, hook)
		}
		re, err := res.R(r, err, http.DefaultClient.Do)
		if err != nil {
			return
		}
		return re.Body, ParseResponse(re, nil)
	}
}

func GetJSON(ctx context.Context, resp any, hook res.Hook[*http.Request]) func(u string) error {
	return func(endpoint string) (err error) {
		r, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if hook != nil {
			r, err = res.R(r, err, hook)
		}
		re, err := res.R(r, err, http.DefaultClient.Do)
		if err != nil {
			return
		}
		defer re.Body.Close()
		err = ParseResponse(re, resp)
		return
	}
}

func Post(ctx context.Context, body io.Reader, resp any, hook res.Hook[*http.Request]) func(u string) (err error) {
	return func(u string) (err error) {
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, u, body)
		if hook != nil {
			r, err = res.R(r, err, hook)
		}
		r, err = res.R(r, err, runnerSetHeaderLine("Content-Type", "application/octet-stream"))
		re, err := res.R(r, err, http.DefaultClient.Do)
		if err != nil {
			return
		}
		defer re.Body.Close()
		return ParseResponse(re, resp)
	}
}

func PostJSON(ctx context.Context, resp any, hook res.Hook[*http.Request]) func(req any) func(u string) (err error) {
	return func(req any) func(u string) (err error) {
		return func(u string) (err error) {
			body, err := JSONReader(req)
			r, err := res.R(body, err, runnerNewRequestWithContext(ctx, http.MethodPost, u))
			if hook != nil {
				r, err = res.R(r, err, hook)
			}
			r, err = res.R(r, err, runnerSetHeaderLine("Content-Type", "application/json"))
			re, err := res.R(r, err, http.DefaultClient.Do)
			if err != nil {
				return
			}
			defer re.Body.Close()
			return ParseResponse(re, resp)
		}
	}
}

func Put(ctx context.Context, body io.Reader, resp any, hook res.Hook[*http.Request]) func(u string) (err error) {
	return func(u string) (err error) {
		r, err := http.NewRequestWithContext(ctx, http.MethodPut, u, body)
		if hook != nil {
			r, err = res.R(r, err, hook)
		}
		r, err = res.R(r, err, runnerSetHeaderLine("Content-Type", "application/octet-stream"))
		re, err := res.R(r, err, http.DefaultClient.Do)
		if err != nil {
			return
		}
		defer re.Body.Close()
		return ParseResponse(re, resp)
	}
}

func Delete(ctx context.Context, hook res.Hook[*http.Request]) func(u string) (err error) {
	return func(u string) (err error) {
		r, err := http.NewRequestWithContext(ctx, http.MethodDelete, u, nil)
		if hook != nil {
			r, err = res.R(r, err, hook)
		}
		re, err := res.R(r, err, http.DefaultClient.Do)
		if err != nil {
			return
		}
		defer re.Body.Close()
		return ParseResponse(re, nil)
	}
}

func runnerNewRequestWithContext(ctx context.Context, method string, u string) func(body io.Reader) (*http.Request, error) {
	return func(body io.Reader) (*http.Request, error) {
		return http.NewRequestWithContext(ctx, method, u, body)
	}
}

func runnerSetHeaderLine(key, value string) func(r *http.Request) (*http.Request, error) {
	return func(r *http.Request) (*http.Request, error) {
		r.Header.Set(key, value)
		return r, nil
	}
}
