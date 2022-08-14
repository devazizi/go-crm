package request

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
