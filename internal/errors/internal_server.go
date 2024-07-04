package errors

import "net/http"

type InternalServerError struct {
	status  int
	message string
}

func (u InternalServerError) GetStatus() int {
	return u.status
}

func (u InternalServerError) GetError() string {
	return u.message
}

func (u InternalServerError) Error() string {
	return u.message
}

func (u InternalServerError) RuntimeError() {
	panic(u.message)
}

func NewInternalServerError(message string) Errors {
	return InternalServerError{status: http.StatusInternalServerError, message: message}
}
