package validation

import (
	"database/sql/driver"
	"reflect"

	"github.com/go-playground/validator/v10"
	"gopkg.in/guregu/null.v4"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()
	v.RegisterCustomTypeFunc(ValidateValuerType, null.String{}, null.Int{}, null.Float{}, null.Bool{})
	return &Validator{
		validator: v,
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func ValidateValuerType(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}
	return nil
}
