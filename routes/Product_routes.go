package routes

import (
	controllers "ecommerce_new/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(r *gin.Engine, db *gorm.DB) {
	productController := controllers.NewProductController(db)

	products := r.Group("/products")
	{
		products.POST("/", productController.CreateProduct)
		products.GET("/:id", productController.GetProduct)
		products.PUT("/:id", productController.UpdateProduct)
		products.DELETE("/:id", productController.DeleteProduct)
		products.GET("/", productController.ListProducts)
	}
}
