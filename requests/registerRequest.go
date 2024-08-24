package requests

type RegisterRequest struct {
	FirstName string `json:"FirstName" binding:"required"`
	LastName string `json:"LastName" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}