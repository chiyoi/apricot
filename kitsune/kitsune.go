package kitsune

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chiyoi/apricot/logs"
)

var ls = func() *logs.Loggers {
	ls := logs.NewLoggers()
	ls.Prefix("[kitsune] ")
	return ls
}()

func SetLogOutput(w io.Writer) {
	ls.SetOutput(w)
}

func ParseRequest(r *http.Request, a any) (err error) {
	if a == nil {
		return nil
	}
	return json.NewDecoder(r.Body).Decode(a)
}

func ParseResponse(re *http.Response, a any) (err error) {
	defer re.Body.Close()

	if re.StatusCode/100 != 2 {
		return newResponseError(re.StatusCode, re.Body)
	}

	if a == nil {
		return
	}
	return json.NewDecoder(re.Body).Decode(&a)
}

type ResponseError struct {
	StatusCode int
	Message    string
}

func newResponseError(code int, body io.Reader) *ResponseError {
	data, err := io.ReadAll(body)
	if err != nil {
		return &ResponseError{
			StatusCode: code,
		}
	}
	return &ResponseError{
		StatusCode: code,
		Message:    string(data),
	}
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("(status: %d, message: %s)", re.StatusCode, re.Message)
}
