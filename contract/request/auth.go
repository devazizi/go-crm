package request

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Paassword string `json:"password"`
}

type ChangePasswordRequest struct {
	CurrentPassword    string `json:"current_password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}
