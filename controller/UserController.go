package controllers

import (
	"ecommerce_new/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := uc.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func (uc *UserController) GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := uc.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)

}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := uc.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := uc.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := uc.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := uc.DB.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

func (uc *UserController) ListUsers(c *gin.Context) {
	var users []models.User

	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

// Implement the other CRUD operations (UpdateUser, DeleteUser, ListUsers)
