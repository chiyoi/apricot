package sakana

import (
	"fmt"
	"io"
)

func InternalError(w io.Writer) {
	if _, err := fmt.Fprintln(w, "Internal error."); err != nil {
		ls.Error("Write output error.", err)
	}
	ls.Panic("internal error")
}

func UsageError(w io.Writer, message string) {
	if _, err := fmt.Fprintln(w, "Usage error:", message); err != nil {
		ls.Error("Write output error.", err)
	}
	ls.Panic("usage error")
}
