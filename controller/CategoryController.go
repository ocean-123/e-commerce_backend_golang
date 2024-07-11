package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ecommerce_new/models"

)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Create(&category).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, category)
}

func (cc *CategoryController) GetCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := cc.DB.First(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(200, category)
}

func (cc *CategoryController) UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := cc.DB.First(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Save(&category).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, category)
}

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := cc.DB.First(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	if err := cc.DB.Delete(&category).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

func (cc *CategoryController) ListCategories(c *gin.Context) {
	var categories []models.Category

	if err := cc.DB.Find(&categories).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, categories)
}
