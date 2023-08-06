package kitsune

import (
	"encoding/json"
	"fmt"
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
	Message    json.RawMessage
}

func newResponseError(code int, reader io.Reader) *ResponseError {
	data, err := io.ReadAll(reader)
	if err != nil {
		return &ResponseError{
			StatusCode: code,
			Message:    []byte(fmt.Sprintf("\"read response failed: %s\"", err.Error())),
		}
	}
	return &ResponseError{
		StatusCode: code,
		Message:    data,
	}
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("(status: %d, message: %s)", re.StatusCode, re.Message)
}
