package response

import "github.com/devazizi/go-crm/entity"

type IndexTaskResponse struct {
	Tasks []entity.Task `json:"tasks"`
}
