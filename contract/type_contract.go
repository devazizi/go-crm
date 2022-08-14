package contract

import (
	"github.com/devazizi/go-crm/contract/request"
	"github.com/devazizi/go-crm/infrastructure"
)

type (
	ValidateLoginRequestFields    func(req request.LoginRequest) error
	ValidateRegisterRequestFields func(registerRequest request.RegisterRequest, DB infrastructure.DB) error
)
