package main

import (
	"ecommerce_new/models"
	"ecommerce_new/routes"
	"ecommerce_new/seed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Order{}, &models.OrderItem{})

	seed.Seed(db)

	// Set up the routes
	routes.SetupUserRoutes(r, db)
	routes.SetupProductRoutes(r, db)
	routes.SetupCategoryRoutes(r, db)
	routes.SetupOrderRoutes(r, db)
	routes.SetupOrderItemRoutes(r, db)

	r.Run()
}

func connectToDatabase() (*gorm.DB, error) {
	dsn := "root:samundra123@tcp(localhost:3306)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
