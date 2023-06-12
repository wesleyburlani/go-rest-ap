package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type ConflictError struct {
	Message string
	Parent  error
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
	}
}

func NewConflictErrorWithParent(message string, parent error) *ConflictError {
	return &ConflictError{
		Message: message,
		Parent:  parent,
	}
}

func IsConflictError(err error) (bool, *ConflictError) {
	var e *ConflictError
	return errors.As(err, &e), e
}

func (e *ConflictError) Error() string {
	if e.Parent != nil {
		return fmt.Sprintf("%s: %s\n", e.Message, e.Parent.Error())
	}
	return e.Message
}

func (e *ConflictError) StatusCode() int {
	return http.StatusConflict
}
