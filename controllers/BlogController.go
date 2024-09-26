package controllers

import (
	"go-authentication/config"
	"go-authentication/models"
	"go-authentication/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strconv"
)

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := c.Get("user_id")
	blog.UserID = userId.(uint)

	if err := config.DB.Create(&blog).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Blog"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog Successfully Created", "data": blog})
}

func ListBlogs(c *gin.Context) {
	var blogs []models.Blog
	config.DB.Preload("User").Preload("Category").Find(&blogs)
	c.JSON(http.StatusOK, gin.H{"message": "Blogs Retrieved", "data": blogs})
}

func GetBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	if err := config.DB.Preload("User").Preload("Category").Preload("Comments").First(&blog, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog Retrieved", "data": blog})
}

func UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	// userId, _ := c.Get("user_id")
	// if blog.UserID != userId.(uint) {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own blogs"})
	// 	return
	// }

	if err := c.ShouldBindJSON(&blog); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	config.DB.Save(&blog)
	c.JSON(http.StatusOK, gin.H{"message": "Blog Updated Successfully", "data": blog})
}

func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	// userID, _ := c.Get("user_id")
    // if blog.UserID != userID.(uint) {
    //     c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own blogs"})
    //     return
    // }

	config.DB.Delete(&blog)
	c.JSON(http.StatusOK, gin.H{"message": "Blog Deleted Successfully"})
}