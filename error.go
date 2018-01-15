package umbrella

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// UnexpectedResponse is returned when the status code is not 200. It contains
// the received status code as well as the body (up to a limit).
type UnexpectedResponse struct {
	Status int
	Body   string
}

func (e UnexpectedResponse) Error() string {
	return fmt.Sprintf("%d %s : %s\n", e.Status, http.StatusText(e.Status), e.Body)
}

// NewUnexpectedResponse reads the first KB of the response body and returns
// a new UnexpectedResponse.
func NewUnexpectedResponse(status int, body io.Reader) UnexpectedResponse {
	b := new(bytes.Buffer)
	io.Copy(b, io.LimitReader(body, 1024))
	return UnexpectedResponse{
		Status: status,
		Body:   strings.TrimSpace(b.String()),
	}
}
