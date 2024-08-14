package errors

import "net/http"

type ConflictError struct {
	Errors
	status  int
	message string
}

func (u ConflictError) GetStatus() int {
	return u.status
}

func (u ConflictError) GetError() string {
	return u.message
}

func (u ConflictError) Error() string {
	return u.message
}

func (u ConflictError) RuntimeError() {
	panic(u.message)
}

func NewConflictError(message string) Errors {
	return ConflictError{status: http.StatusConflict, message: message}
}
