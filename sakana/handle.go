package sakana

import (
	"fmt"
	"io"
)

func InternalError(w io.Writer, message string) {
	var msg string
	if message != "" {
		msg = "Internal error: " + message
	} else {
		msg = "Internal error."
	}
	if _, err := fmt.Fprintln(w, msg); err != nil {
		ls.Error("Write output error.", err)
	}
}

func UsageError(w io.Writer, message string) {
	var msg string
	if message != "" {
		msg = "Usage error: " + message
	} else {
		msg = "Usage error."
	}
	if _, err := fmt.Fprintln(w, msg); err != nil {
		ls.Error("Write output error.", err)
	}
}
