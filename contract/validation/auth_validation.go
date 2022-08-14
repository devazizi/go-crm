package validation

import (
	"github.com/devazizi/go-crm/contract/request"
	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
)

func ValidateLoginRequestFields(req request.LoginRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(6, 256)),
	)
}
