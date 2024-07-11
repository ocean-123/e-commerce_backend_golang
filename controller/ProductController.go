package controllers

import (
	"ecommerce_new/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{DB: db}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, product)
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := pc.DB.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := pc.DB.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := pc.DB.Save(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := pc.DB.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	if err := pc.DB.Delete(&product).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

func (pc *ProductController) ListProducts(c *gin.Context) {
	var products []models.Product

	if err := pc.DB.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}
