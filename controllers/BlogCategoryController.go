package controllers

import (
	"go-authentication/config"
	"go-authentication/models"
	"go-authentication/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the model struct
	if err := validate.Struct(&category); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		utils.Logger.Printf("Validation Error: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": validationErrors.Error()})
		return
	}

	// Start the DB Transaction
	tx := config.DB.Begin()
	if tx.Error != nil {
		utils.Logger.Printf("Transaction Error: %v", tx.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	if err := tx.Create(&category).Error; err != nil {
		// Rollback the transaction in case of error
		tx.Rollback()
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog Category Successfully Created", "data": category})
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Category Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category Retrieved", "data": category})
}

func ListCategories(c *gin.Context) {
	var categories []models.Category
	config.DB.Preload("Blog").Find(&categories)
	c.JSON(http.StatusOK, gin.H{"message": "Categories Retrieved", "data": categories})
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	config.DB.Save(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Category Updated Successfully", "data": category})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		utils.Logger.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	config.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"message": "Category Deleted Successfully"})
}