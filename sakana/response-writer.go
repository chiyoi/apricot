package sakana

import (
	"io"
	"sync"
)

type ResponseWriter interface {
	io.Writer
	Out() io.Writer
	Err() io.Writer
}

func NewResponseWriter(out, err io.Writer) ResponseWriter {
	return &responseWriter{out: out, err: err}
}

type responseWriter struct {
	mu       sync.Mutex
	out, err io.Writer
}

var _ ResponseWriter = &responseWriter{}

func (w *responseWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.out.Write(p)
}

func (w *responseWriter) Out() io.Writer {
	return w.out
}

func (w *responseWriter) Err() io.Writer {
	return w.err
}
