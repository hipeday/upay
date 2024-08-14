package errors

import "net/http"

type NotFoundError struct {
	Errors
	status  int
	message string
}

func (u NotFoundError) GetStatus() int {
	return u.status
}

func (u NotFoundError) GetError() string {
	return u.message
}

func (u NotFoundError) Error() string {
	return u.message
}

func (u NotFoundError) RuntimeError() {
	panic(u.message)
}

func NewNotFoundError(message string) Errors {
	return NotFoundError{status: http.StatusNotFound, message: message}
}
