package kitsune

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chiyoi/iter/res"
)

func ExampleThreadsIDMessagesGet(ctx context.Context, threadID string, limit int, before, after string) (messages []Message, err error) {
	endpoint := "https://example.com"
	setAuth := func(r *http.Request) (*http.Request, error) {
		r.Header.Set("Authorization", "Bearer "+"<my token>")
		return r, nil
	}

	setQuery := func(r *http.Request) (*http.Request, error) {
		q := r.URL.Query()
		q.Set("limit", strconv.Itoa(limit))
		q.Set("before", before)
		q.Set("after", after)
		r.URL.RawQuery = q.Encode()
		return r, nil
	}
	u, err := url.JoinPath(endpoint, "threads", threadID, "messages")

	var resp struct {
		Messages []string `json:"messages"`
	}
	_, err = res.Then(u, err, GetJSON(ctx, &resp, func(r *http.Request) (*http.Request, error) {
		var err error
		r, err = res.Then(r, err, setAuth)
		r, err = res.Then(r, err, setQuery)
		return r, err
	}))
	messages = resp.Messages
	return
}
