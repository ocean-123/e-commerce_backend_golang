package routes

import (
	"ecommerce_new/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	users := r.Group("/users")
	{
		users.POST("/", userController.CreateUser)
		users.GET("/:id", userController.GetUser)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
		users.GET("/", userController.ListUsers)
	}
}
