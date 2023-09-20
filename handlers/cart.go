package handlers

import (
	"eshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler struct {
	db *gorm.DB
}

func NewCartHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{db}
}

func (h *CartHandler) CreateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateCart(h.db, &cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cart)
}

func (h *CartHandler) GetCartByID(c *gin.Context) {
	cartID := c.Param("id")

	cart, err := models.GetCartByID(h.db, cartID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) UpdateCart(c *gin.Context) {
	cartID := c.Param("id")

	var updateCart models.Cart
	if err := c.ShouldBindJSON(&updateCart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the cart exists
	cart, err := models.GetCartByID(h.db, cartID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Update the cart fields as needed
	// ...

	if err := models.UpdateCart(h.db, cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *CartHandler) DeleteCart(c *gin.Context) {
	cartID := c.Param("id")

	cart, err := models.GetCartByID(h.db, cartID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	if err := models.DeleteCart(h.db, cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
