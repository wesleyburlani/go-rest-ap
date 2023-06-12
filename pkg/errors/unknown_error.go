package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type UnknownError struct {
	Message string
	Parent  error
}

func NewUnknownError(message string) *UnknownError {
	return &UnknownError{
		Message: message,
	}
}

func NewUnknownErrorWithParent(message string, parent error) *UnknownError {
	return &UnknownError{
		Message: message,
		Parent:  parent,
	}
}

func IsUnknownError(err error) (bool, *UnknownError) {
	var e *UnknownError
	return errors.As(err, &e), e
}

func (e *UnknownError) Error() string {
	if e.Parent != nil {
		return fmt.Sprintf("%s: %s\n", e.Message, e.Parent.Error())
	}
	return e.Message
}

func (e *UnknownError) StatusCode() int {
	return http.StatusInternalServerError
}
