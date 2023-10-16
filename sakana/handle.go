package sakana

import (
	"fmt"
	"io"
)

func InternalError(w io.Writer) {
	fmt.Fprintln(w, "Internal error.")
}

func UsageError(w io.Writer, message string, usage string) {
	if message != "" {
		fmt.Fprintln(w, "Usage error:", message)
	} else {
		fmt.Fprintln(w, "Usage error.")
	}
	if usage != "" {
		fmt.Fprintln(w, usage)
	}
}
