package neko

import (
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/chiyoi/apricot/test"
)

func TestStatusHandler(t *testing.T) {
	tcs := [][2]string{
		{"v0.1.0", "Nyan~\nVersion: v0.1.0\n"},
		{"aaa", "Nyan~\nVersion: aaa\n"},
	}
	h := StatusHandler()
	for _, tc := range tcs {
		in, out := tc[0], tc[1]
		os.Setenv("VERSION", in)

		var buf test.ResponseBuffer
		h.ServeHTTP(&buf, &http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme:  "http",
				Host:    "localhost",
				RawPath: "/",
			},
		})
		test.AssertEqual(t, out, buf.Body.String())
	}
}
