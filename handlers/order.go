package handlers

import (
	"eshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OrderHandler handles HTTP requests for the Order model.
type OrderHandler struct {
	db *gorm.DB
}

// NewOrderHandler creates a new OrderHandler with the provided database connection.
func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{db}
}

// CreateOrder handles the creation of a new order.
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateOrder(h.db, &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrderByID retrieves an order by its ID.
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	order, err := models.GetOrderByID(h.db, orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// UpdateOrder updates an existing order.
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	orderID := c.Param("id")
	var updatedOrder models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := models.GetOrderByID(h.db, orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Update order fields here
	order.TotalPrice = updatedOrder.TotalPrice
	order.OrderStatus = updatedOrder.OrderStatus
	order.PaymentStatus = updatedOrder.PaymentStatus

	if err := models.UpdateOrder(h.db, order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DeleteOrder deletes an order.
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	order, err := models.GetOrderByID(h.db, orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := models.DeleteOrder(h.db, order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
