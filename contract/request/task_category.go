package request

type CreateTaskCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateTaskCategoryRequest struct {
	Name string `json:"name"`
}
