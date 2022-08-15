package validation

import (
	"github.com/devazizi/go-crm/contract/request"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateCreateTaskRequest(req request.CreateTaskRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 256)),
		validation.Field(&req.CategoryID, validation.Required),
		validation.Field(&req.AssignTo, validation.Required),
	)
}
