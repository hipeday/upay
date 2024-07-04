package errors

import "net/http"

type IllegalArgumentError struct {
	status  int
	message string
}

func (u IllegalArgumentError) GetStatus() int {
	return u.status
}

func (u IllegalArgumentError) GetError() string {
	return u.message
}

func (u IllegalArgumentError) Error() string {
	return u.message
}

func (u IllegalArgumentError) RuntimeError() {
	panic(u.message)
}

func NewIllegalArgumentError(message string) Errors {
	return IllegalArgumentError{status: http.StatusBadRequest, message: message}
}
