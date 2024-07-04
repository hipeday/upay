package errors

type Errors interface {
	GetStatus() int
	GetError() string
	error
}
