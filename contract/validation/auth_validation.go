package validation

import (
	"github.com/devazizi/go-crm/contract/request"
	"github.com/devazizi/go-crm/contract/validation/rules"
	"github.com/devazizi/go-crm/infrastructure"
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

func ValidateRegisterRequestFields(registerRequest request.RegisterRequest, DB infrastructure.DB) error {

	return validation.ValidateStruct(
		&registerRequest,
		validation.Field(&registerRequest.Name, validation.Required, validation.Length(3, 256)),
		validation.Field(&registerRequest.Paassword, validation.Required, validation.Length(6, 256)),
		validation.Field(&registerRequest.Email, validation.Required, is.Email, validation.By(rules.CheckEmailMustUnique(DB))),
	)
}
