package handlers

import (
	"eshop/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductHandler is a struct that holds the database connection.
type ProducItemtHandler struct {
	db *gorm.DB
}

// NewProductHandler creates a new ProductHandler with the provided database connection.
func NewProductItemHandler(db *gorm.DB) *ProducItemtHandler {
	return &ProducItemtHandler{db}
}

func (h *ProducItemtHandler) CreateProductItem(c *gin.Context) {

	fmt.Println("we are here at productitem creation")
	// storeID := c.Param("store_id")
	productId := c.Param("product_id")

	// Parse JSON input and validate it
	var productItem models.ProductItem
	if err := c.ShouldBindJSON(&productItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// product.ID = utils.NewRandomUUID(c)
	productItem.ProductID = productId

	// check if category exists in the store
	// this is done by the gorm Before save function

	// Insert the product into the database
	if err := models.CreateProductItem(h.db, &productItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, productItem)
}
