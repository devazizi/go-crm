package response

const SUCCESSFUL_MESSAGE = "your action successful"

const SUCCESS = true

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
