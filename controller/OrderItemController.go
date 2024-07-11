package controllers

import (
	"ecommerce_new/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderItemController struct {
	DB *gorm.DB
}

func NewOrderItemController(db *gorm.DB) *OrderItemController {
	return &OrderItemController{DB: db}
}

func (oic *OrderItemController) CreateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := oic.DB.Create(&orderItem).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, orderItem)
}

func (oic *OrderItemController) GetOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	id := c.Param("id")

	if err := oic.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order Item not found"})
		return
	}

	c.JSON(200, orderItem)
}

func (oic *OrderItemController) UpdateOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	id := c.Param("id")

	if err := oic.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := oic.DB.Save(&orderItem).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orderItem)
}

func (oic *OrderItemController) DeleteOrderItem(c *gin.Context) {
	var orderItem models.OrderItem
	id := c.Param("id")

	if err := oic.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Order Item not found"})
		return
	}

	if err := oic.DB.Delete(&orderItem).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, nil)
}

func (oic *OrderItemController) ListOrderItems(c *gin.Context) {
	var orderItems []models.OrderItem

	if err := oic.DB.Find(&orderItems).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orderItems)
}
