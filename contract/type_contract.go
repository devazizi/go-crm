package contract

import "github.com/devazizi/go-crm/contract/request"

type (
	ValidateLoginRequestFields func(req request.LoginRequest) error
)
