package request

type CreateTaskRequest struct {
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	AssignTo    []int  `json:"assign_to"`
	Description string `json:"description"`
}
