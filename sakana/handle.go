package sakana

import (
	"fmt"
	"io"
)

func InternalError(w io.Writer) {
	fmt.Fprintln(w, "Internal error.")
}

func UsageError(w io.Writer, message string, usage string) {
	fmt.Fprintln(w, "Usage error.")
	fmt.Fprintln(w, message)
	if usage != "" {
		fmt.Fprintln(w, usage)
	}
}
