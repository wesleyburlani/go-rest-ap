package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type NotFoundError struct {
	Message string
	Parent  error
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
	}
}

func NewNotFoundErrorWithParent(message string, parent error) *NotFoundError {
	return &NotFoundError{
		Message: message,
		Parent:  parent,
	}
}

func IsNotFoundError(err error) (bool, *NotFoundError) {
	var e *NotFoundError
	return errors.As(err, &e), e
}

func (e *NotFoundError) Error() string {
	if e.Parent != nil {
		return fmt.Sprintf("%s: %s\n", e.Message, e.Parent.Error())
	}
	return e.Message
}

func (e *NotFoundError) StatusCode() int {
	return http.StatusNotFound
}
