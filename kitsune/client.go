package kitsune

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get(endpoint string) (r io.ReadCloser, err error) {
	return Client{Context: context.Background(), Endpoint: endpoint}.Get()
}

func Put(endpoint string, body io.Reader) (r io.ReadCloser, err error) {
	return Client{Context: context.Background(), Endpoint: endpoint}.Put(body)
}

func Post(endpoint string, req any) (r io.ReadCloser, err error) {
	return Client{Context: context.Background(), Endpoint: endpoint}.Post(req)
}

func Delete(endpoint string) (r io.ReadCloser, err error) {
	return Client{Context: context.Background(), Endpoint: endpoint}.Delete()
}

type Client struct {
	Context  context.Context
	Endpoint string
	Header   http.Header
}

func (c Client) Get() (r io.ReadCloser, err error) {
	hr, err := http.NewRequestWithContext(c.Context, http.MethodGet, c.Endpoint, nil)
	if err != nil {
		return
	}

	for k := range c.Header {
		hr.Header.Set(k, c.Header.Get(k))
	}

	hre, err := http.DefaultClient.Do(hr)
	if err != nil {
		return
	}

	if hre.StatusCode != http.StatusOK {
		return nil, newResponseError(hre.StatusCode, hre.Body)
	}
	return hre.Body, nil
}

func (c Client) Put(body io.Reader) (r io.ReadCloser, err error) {
	hr, err := http.NewRequestWithContext(c.Context, http.MethodPut, c.Endpoint, body)
	if err != nil {
		return
	}

	hr.Header.Set("Content-Type", "application/octet-stream")
	for k := range c.Header {
		hr.Header.Set(k, c.Header.Get(k))
	}

	hre, err := http.DefaultClient.Do(hr)
	if err != nil {
		return
	}

	if hre.StatusCode != http.StatusOK {
		return nil, newResponseError(hre.StatusCode, hre.Body)
	}
	return hre.Body, nil
}

func (c Client) Post(req any) (r io.ReadCloser, err error) {
	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(req); err != nil {
		return
	}

	hr, err := http.NewRequestWithContext(c.Context, http.MethodPost, c.Endpoint, &buf)
	if err != nil {
		return
	}

	hr.Header.Set("Content-Type", "application/json")

	for k := range c.Header {
		hr.Header.Set(k, c.Header.Get(k))
	}

	hre, err := http.DefaultClient.Do(hr)
	if err != nil {
		return
	}

	if hre.StatusCode != http.StatusOK {
		return nil, newResponseError(hre.StatusCode, hre.Body)
	}
	return hre.Body, nil
}

func (c Client) Delete() (r io.ReadCloser, err error) {
	hr, err := http.NewRequestWithContext(c.Context, http.MethodDelete, c.Endpoint, nil)
	if err != nil {
		return
	}

	for k := range c.Header {
		hr.Header.Set(k, c.Header.Get(k))
	}

	hre, err := http.DefaultClient.Do(hr)
	if err != nil {
		return
	}

	if hre.StatusCode != http.StatusOK {
		return nil, newResponseError(hre.StatusCode, hre.Body)
	}
	return hre.Body, nil
}

func ParseResponse(r io.ReadCloser, resp any) (err error) {
	defer r.Close()

	if resp == nil {
		return
	}
	return json.NewDecoder(r).Decode(&resp)
}

type ResponseError struct {
	StatusCode int
	Message    json.RawMessage
}

func newResponseError(code int, reader io.Reader) *ResponseError {
	data, err := io.ReadAll(reader)
	if err != nil {
		return &ResponseError{
			StatusCode: code,
			Message:    []byte(fmt.Sprintf("\"read response failed: %s\"", err.Error())),
		}
	}
	return &ResponseError{
		StatusCode: code,
		Message:    data,
	}
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("(status: %d, message: %s)", re.StatusCode, re.Message)
}
