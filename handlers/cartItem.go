package handlers

import (
	"eshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartItemHandler struct {
	db *gorm.DB
}

func NewCartItemHandler(db *gorm.DB) *CartItemHandler {
	return &CartItemHandler{db}
}

func (h *CartItemHandler) CreateCartItem(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateCartItem(h.db, &cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cartItem)
}

func (h *CartItemHandler) GetCartItemByID(c *gin.Context) {
	cartItemID := c.Param("id")

	cartItem, err := models.GetCartItemByID(h.db, cartItemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	c.JSON(http.StatusOK, cartItem)
}

func (h *CartItemHandler) UpdateCartItem(c *gin.Context) {
	cartItemID := c.Param("id")

	var updateCartItem models.CartItem
	if err := c.ShouldBindJSON(&updateCartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the cart item exists
	cartItem, err := models.GetCartItemByID(h.db, cartItemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	// Update the cart item fields as needed
	// ...

	if err := models.UpdateCartItem(h.db, cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartItem)
}

func (h *CartItemHandler) DeleteCartItem(c *gin.Context) {
	cartItemID := c.Param("id")

	cartItem, err := models.GetCartItemByID(h.db, cartItemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	if err := models.DeleteCartItem(h.db, cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
