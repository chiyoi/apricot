package sakana

import (
	"fmt"
	"io"
)

func InternalError(w io.Writer) {
	if _, err := fmt.Fprintln(w, "Internal error."); err != nil {
		ls.Error(err)
	}
}

func UsageError(w io.Writer, message string, usage string) {
	if _, err := fmt.Fprintln(w, "Usage error."); err != nil {
		ls.Error(err)
		return
	}
	if _, err := fmt.Fprintln(w, message); err != nil {
		ls.Error(err)
		return
	}
	if usage != "" {
		if _, err := fmt.Fprintln(w, usage); err != nil {
			ls.Error(err)
			return
		}
	}
}
