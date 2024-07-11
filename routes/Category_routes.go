package routes

import (
	"ecommerce_new/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCategoryRoutes(r *gin.Engine, db *gorm.DB) {
	categoryController := controllers.NewCategoryController(db)

	categories := r.Group("/categories")
	{
		categories.POST("/", categoryController.CreateCategory)
		categories.GET("/:id", categoryController.GetCategory)
		categories.PUT("/:id", categoryController.UpdateCategory)
		categories.DELETE("/:id", categoryController.DeleteCategory)
		categories.GET("/", categoryController.ListCategories)
	}
}
