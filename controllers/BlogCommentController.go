package controllers

import (
	"go-authentication/config"
	"go-authentication/models"
	"go-authentication/utils"
	"net/http"
	// "strconv"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
    comment.UserID = userID.(uint)

	// blogID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    // if err != nil {
    //     c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
    //     return
    // }
    // comment.BlogID = uint(blogID)

    if err := config.DB.Create(&comment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
        return
    }

	c.JSON(http.StatusCreated, gin.H{"message": "Blog Comment Successfully Created", "data": comment})
}

func GetComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment Retrieved", "data": comment})
}

func ListComments(c *gin.Context) {
	var comments []models.Comment
	config.DB.Preload("Blog").Preload("User").Find(&comments)
	c.JSON(http.StatusOK, gin.H{"message": "Comments Retrieved", "data": comments})
}

func UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	config.DB.Save(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Comment Updated Successfully", "data": comment})
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	config.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "Comment Deleted Successfully"})
}