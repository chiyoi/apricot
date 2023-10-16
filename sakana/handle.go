package sakana

import (
	"fmt"
)

func InternalError(w ResponseWriter, message string) {
	var msg string
	if message != "" {
		msg = "Internal error: " + message
	} else {
		msg = "Internal error."
	}
	if _, err := fmt.Fprintln(w.Err(), msg); err != nil {
		ls.Error("Write output error.", err)
	}
}

func UsageError(w ResponseWriter, message string) {
	var msg string
	if message != "" {
		msg = "Usage error: " + message
	} else {
		msg = "Usage error."
	}
	if _, err := fmt.Fprintln(w.Err(), msg); err != nil {
		ls.Error("Write output error.", err)
	}
}
