package errors

import (
	"fmt"
)

type AuthenticationError struct {
	StatusCode int
	Err        error
	Message    string
}

func NewAuthenticationError(statusCode int, err error, message string) *AuthenticationError {
	return &AuthenticationError{statusCode, err, message}

}
func (r *AuthenticationError) Error() string {
	return fmt.Sprintf("%d\n , %v\n , %s\n", r.StatusCode, r.Err, r.Message)
}
