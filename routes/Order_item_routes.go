package routes

import (
	"ecommerce_new/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrderItemRoutes(r *gin.Engine, db *gorm.DB) {
	orderItemController := controllers.NewOrderItemController(db)

	orderItems := r.Group("/order-items")
	{
		orderItems.POST("/", orderItemController.CreateOrderItem)
		orderItems.GET("/:id", orderItemController.GetOrderItem)
		orderItems.PUT("/:id", orderItemController.UpdateOrderItem)
		orderItems.DELETE("/:id", orderItemController.DeleteOrderItem)
		orderItems.GET("/", orderItemController.ListOrderItems)
	}
}
