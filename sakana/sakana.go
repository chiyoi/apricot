package sakana

import (
	"io"

	"github.com/chiyoi/apricot/logs"
)

var ls = func() (ls *logs.Loggers) {
	ls = logs.NewLoggers()
	ls.Prefix("[sakana]")
	return
}()

func SetLogOutput(w io.Writer) {
	ls.SetOutput(w)
}

func SetLogLevel(l logs.Level) {
	ls.SetLevel(l)
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
