package sakana

import (
	"fmt"
	"io"
	"testing"
)

type tc struct {
	args []string
}

var tcs = []tc{
	{[]string{"--help"}},
	// {[]string{"--nyan"}},
	// {[]string{}},
	// {[]string{"nyan"}},
}

func TestCommand(t *testing.T) {
	c := NewCommand("neko")
	c.Summary("neko", "Nyan~")
	c.Command("c1", "d1", nil)
	c.Command("c2", "d2", nil)
	c.Command("c3", "d3", nil)

	c.Work(HandlerFunc(func(w io.Writer, args []string) {
		if _, err := fmt.Fprintln(w, "nyan~"); err != nil {
			t.Error(err)
		}
	}))

	for _, tc := range tcs {
		t.Log("input:", tc.args)
		func() {
			defer func() {
				if err := recover(); err != nil {
					t.Log("panicked:", err)
				}
			}()
			c.ServeArgs(nil, tc.args)
		}()
	}
}
