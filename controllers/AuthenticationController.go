package controllers

import (
	"go-authentication/config"
	"go-authentication/models"
	"go-authentication/utils"
	"go-authentication/requests"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Hash Password"})
		return
	}
	user.Password = string(hashedPassword)

	// Save User to Database
	if err := config.DB.Create(&user).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Register User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var loginRequest requests.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

func ForgotPassword(c *gin.Context) {
	var forgotPasswordRequest requests.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&forgotPasswordRequest); err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", forgotPasswordRequest.Email).First(&user).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Generate reset token
	resetToken := utils.GenerateResetToken()
	user.ResetToken = resetToken
	expiryTime := time.Now().Add(15 * time.Minute)
    user.ResetTokenExpiry = &expiryTime

	if err := config.DB.Save(&user).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save reset token"})
		return
	}

	// TODO: Send reset token to user's email

	c.JSON(http.StatusOK, gin.H{"message": "Reset token sent to email"})
}

func ResetPassword(c *gin.Context) {
	var resetPasswordRequest requests.ResetPasswordRequest
	if err := c.ShouldBindJSON(&resetPasswordRequest); err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("reset_token = ?", resetPasswordRequest.ResetToken).First(&user).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid reset token"})
		return
	}

	// Check if token is expired
	if user.ResetTokenExpiry.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reset token has expired"})
		return
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resetPasswordRequest.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)
	user.ResetToken = ""
	user.ResetTokenExpiry = nil

	if err := config.DB.Save(&user).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}