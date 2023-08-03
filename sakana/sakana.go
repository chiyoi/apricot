package sakana

import (
	"io"
	"os"

	"github.com/chiyoi/apricot/logs"
)

var ls = logs.NewLoggers()
var output io.Writer = os.Stdout

func init() {
	ls.PrependPrefix("[sakana] ")
}

func SetLogFile(w io.Writer) { ls.SetOutput(w) }

type Handler interface {
	ServeArgs(w io.Writer, args []string)
}

type HandlerFunc func(w io.Writer, args []string)

func (h HandlerFunc) ServeArgs(w io.Writer, args []string) { h(w, args) }
