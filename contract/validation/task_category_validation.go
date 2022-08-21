package validation

import (
	"github.com/devazizi/go-crm/contract/request"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateCreateTaskCategoryRequest(req request.CreateTaskCategoryRequest) error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 256)),
	)
}
