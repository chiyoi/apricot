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
	ServeArgs(w ResponseWriter, args []string) int
}

type HandlerFunc func(w ResponseWriter, args []string) int

var _ Handler = (HandlerFunc)(nil)

func (h HandlerFunc) ServeArgs(w ResponseWriter, args []string) int {
	return h(w, args)
}
