package response

import "github.com/devazizi/go-crm/entity"

type LoginResponse struct {
	User  entity.User `json:"user"`
	Token string      `json:"token"`
}
