package test

import (
	"bytes"
	"net/http"
)

type ResponseBuffer struct {
	StatusCode int
	Body       bytes.Buffer

	h http.Header
}

func (w *ResponseBuffer) Write(bs []byte) (int, error) {
	return w.Body.Write(bs)
}

func (w *ResponseBuffer) Header() http.Header {
	if w.h == nil {
		w.h = http.Header{}
	}
	return w.h
}

func (w *ResponseBuffer) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
}
