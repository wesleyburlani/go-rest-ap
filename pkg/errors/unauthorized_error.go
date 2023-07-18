package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type UnauthorizedError struct {
	Message string
	Parent  error
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Message: message,
	}
}

func NewUnauthorizedErrorWithParent(message string, parent error) *UnauthorizedError {
	return &UnauthorizedError{
		Message: message,
		Parent:  parent,
	}
}

func IsUnauthorizedError(err error) (bool, *UnauthorizedError) {
	var e *UnauthorizedError
	return errors.As(err, &e), e
}

func (e *UnauthorizedError) Error() string {
	if e.Parent != nil {
		return fmt.Sprintf("%s: %s\n", e.Message, e.Parent.Error())
	}
	return e.Message
}

func (e *UnauthorizedError) StatusCode() int {
	return http.StatusUnauthorized
}
