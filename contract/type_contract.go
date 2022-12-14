package contract

import (
	"github.com/devazizi/go-crm/contract/request"
	"github.com/devazizi/go-crm/infrastructure"
)

type (
	ValidateLoginRequestFields          func(req request.LoginRequest) error
	ValidateRegisterRequestFields       func(registerRequest request.RegisterRequest, DB infrastructure.DB) error
	ValidateCreateTaskRequest           func(req request.CreateTaskRequest) error
	ValidateCreateTaskCategoryRequest   func(req request.CreateTaskCategoryRequest) error
	ValidateUpdateTaskCategoryRequest   func(req request.UpdateTaskCategoryRequest) error
	ValidateChangePasswordRequestFields func(request request.ChangePasswordRequest) error
	ValidateForgetPasswordRequestFields func(request request.ForgetPasswordRequest) error
)
