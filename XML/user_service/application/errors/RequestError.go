package errors

import (
	"fmt"
)

type RequestError struct {
	StatusCode int
	Err        error
	Message    string
}

func NewRequestError(statusCode int, err error, message string) *RequestError {
	return &RequestError{statusCode, err, message}

}
func (r *RequestError) Error() string {
	return fmt.Sprintf("\n Costume error : "+
		"Status:  %d \n"+
		"Err: %v \n"+
		"Message: %s \n",
		r.StatusCode, r.Err, r.Message)
}
