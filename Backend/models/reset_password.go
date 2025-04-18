package models

type ForgotPasswordRequest struct {
    Email string `json:"email" binding:"required"`
}
 
type ResetPasswordRequest struct {
	Email  string `json:"email"`
	NewPassword string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
