package errors

type HttpError interface {
	Error() string
	StatusCode() int
}
