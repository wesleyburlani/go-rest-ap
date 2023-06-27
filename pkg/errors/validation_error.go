package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type ValidationError struct {
	Message string
	Parent  error
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		Message: message,
	}
}

func NewValidationErrorWithParent(message string, parent error) *ValidationError {
	return &ValidationError{
		Message: message,
		Parent:  parent,
	}
}

func IsValidationError(err error) (bool, *ValidationError) {
	var e *ValidationError
	return errors.As(err, &e), e
}

func (e *ValidationError) Error() string {
	if e.Parent != nil {
		return fmt.Sprintf("%s: %s\n", e.Message, e.Parent.Error())
	}
	return e.Message
}

func (e *ValidationError) StatusCode() int {
	return http.StatusBadRequest
}
