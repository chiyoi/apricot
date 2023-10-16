package sakana

import (
	"io"

	"github.com/chiyoi/apricot/logs"
)

var ls = logs.NewLoggers()

func init() {
	ls.Prefix("[sakana] ")
}

func SetLogFile(w io.Writer) {
	ls.SetOutput(w)
}

type Handler interface {
	ServeArgs(out [2]io.Writer, in io.Reader, args ...string) int
}

type HandlerFunc func(out [2]io.Writer, in io.Reader, args ...string) int

var _ Handler = (HandlerFunc)(nil)

func (h HandlerFunc) ServeArgs(out [2]io.Writer, in io.Reader, args ...string) int {
	return h(out, in, args...)
}
