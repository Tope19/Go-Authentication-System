package requests

type ResetPasswordRequest struct {
	ResetToken string `json:"reset_token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}