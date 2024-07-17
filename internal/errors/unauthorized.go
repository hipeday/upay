package errors

import "net/http"

type UnauthorizedError struct {
	status  int
	message string
}

func (u UnauthorizedError) GetStatus() int {
	return u.status
}

func (u UnauthorizedError) GetError() string {
	return u.message
}

func (u UnauthorizedError) Error() string {
	return u.message
}

func (u UnauthorizedError) RuntimeError() {
	panic(u.message)
}

func NewUnauthorizedError(message string) Errors {
	return UnauthorizedError{status: http.StatusUnauthorized, message: message}
}
