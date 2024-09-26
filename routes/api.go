package routes

import (
	"go-authentication/controllers"
	"go-authentication/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Group Route
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
			auth.POST("/forgot-password", controllers.ForgotPassword)
			auth.POST("/reset-password", controllers.ResetPassword)
		}

		// Authenticated Protected Routes
		blog := v1.Group("/blogs")
		blog.Use(middleware.AuthMiddleware())
		{
			blog.POST("/create", controllers.CreateBlog)
			blog.GET("/", controllers.ListBlogs)
			blog.GET("/:id", controllers.GetBlog)
			blog.PUT("/:id", controllers.UpdateBlog)
			blog.DELETE("/:id", controllers.DeleteBlog)
		}

		category := v1.Group("/categories")
		category.Use(middleware.AuthMiddleware())
		{
			category.POST("/create", controllers.CreateCategory)
			category.GET("/", controllers.ListCategories)
			category.GET("/:id", controllers.GetCategory)
			category.PUT("/:id", controllers.UpdateCategory)
			category.DELETE("/:id", controllers.DeleteCategory)
		}

		comment := v1.Group("/comments")
		comment.Use(middleware.AuthMiddleware())
		{
			comment.POST("/create", controllers.CreateComment)
			comment.GET("/", controllers.ListComments)
			comment.GET("/:id", controllers.GetComment)
			comment.PUT("/:id", controllers.UpdateComment)
			comment.DELETE("/:id", controllers.DeleteComment)
		}
	}
}