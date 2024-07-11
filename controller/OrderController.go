package controllers

import (
	"ecommerce_new/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{DB: db}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := oc.DB.Create(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, order)
}

func (oc *OrderController) GetOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := oc.DB.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, order)
}

func (oc *OrderController) UpdateOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := oc.DB.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := oc.DB.Save(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, order)
}

func (oc *OrderController) DeleteOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := oc.DB.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	if err := oc.DB.Delete(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

func (oc *OrderController) ListOrders(c *gin.Context) {
	var orders []models.Order

	if err := oc.DB.Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}
