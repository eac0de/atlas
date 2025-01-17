package httperror

import (
	"errors"
	"net/http"
)

type HTTPError struct {
	err        error
	msg        string
	statusCode int
}

func (e *HTTPError) Error() string {
	return e.msg
}
func (e *HTTPError) Unwrap() error {
	return e.err
}

func New(err error, msg string, statusCode int) *HTTPError {
	return &HTTPError{
		err:        err,
		msg:        msg,
		statusCode: statusCode,
	}
}

func GetMessageAndStatusCode(err error) (string, int) {
	var e *HTTPError
	if errors.As(err, &e) {
		return e.msg, e.statusCode
	} else {
		return err.Error(), http.StatusInternalServerError
	}
}

func IsNotFound(err error) bool {
	var e *HTTPError
	if errors.As(err, &e) {
		if e.statusCode == http.StatusNotFound {
			return true
		}
	}
	return false
}
