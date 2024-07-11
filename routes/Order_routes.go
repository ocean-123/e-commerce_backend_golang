package routes

import (
	"ecommerce_new/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrderRoutes(r *gin.Engine, db *gorm.DB) {
	orderController := controllers.NewOrderController(db)

	orders := r.Group("/orders")
	{
		orders.POST("/", orderController.CreateOrder)
		orders.GET("/:id", orderController.GetOrder)
		orders.PUT("/:id", orderController.UpdateOrder)
		orders.DELETE("/:id", orderController.DeleteOrder)
		orders.GET("/", orderController.ListOrders)
	}
}
