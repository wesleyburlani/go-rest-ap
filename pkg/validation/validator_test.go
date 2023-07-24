package validation_test

import (
	"testing"

	"github.com/wesleyburlani/go-rest-api/pkg/validation"
	"gopkg.in/guregu/null.v4"
)

type OptionalEmailStruct struct {
	Email null.String `json:"email" validate:"omitempty,email"`
}

func TestOptionalEmailStruct_NullEmail(t *testing.T) {
	v := validation.NewValidator()
	email := OptionalEmailStruct{
		Email: null.StringFromPtr(nil),
	}
	err := v.Validate(email)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestOptionalEmailStruct_EmptyEmail(t *testing.T) {
	v := validation.NewValidator()
	email := OptionalEmailStruct{
		Email: null.StringFromPtr(new(string)),
	}
	err := v.Validate(email)
	if err != nil {
		t.Errorf("expected no error, got nil")
	}
}
