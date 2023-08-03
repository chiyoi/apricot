package kitsune

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/chiyoi/apricot/logs"
)

var ls = logs.NewLoggers()

func init() {
	ls.PrependPrefix("[kitsune] ")
}

func SetLogFile(w io.Writer) {
	ls.SetOutput(w)
}

func ParseRequest(w http.ResponseWriter, r *http.Request, req any) (ok bool) {
	if req == nil {
		return true
	}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		message := "Parse request error."
		ls.Warning(message, err)
		BadRequest(w, message)
		return
	}
	return true
}
