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
	ServeArgs(f Files, args []string) int
}

type Files struct {
	In       io.Reader
	Out, Err io.Writer
}

type HandlerFunc func(f Files, args []string) int

var _ Handler = (HandlerFunc)(nil)

func (h HandlerFunc) ServeArgs(f Files, args []string) int {
	return h(f, args)
}
