package handlers

import (
	"errors"
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
func NewProductItemHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{db}
}

func (h *ProductHandler) CreateProductItem(c *gin.Context) {
	storeID := c.Param("store_id")
	// productId := c.Param("product_id")
	userID, _ := c.Get("user_id") // not using exists here since if user_id doesn't exists then it shouldn't pass the middleware

	// check if the store exists
	store, err := models.GetStoreByID(h.db, storeID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Store doesn't exist",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	// check if the user is the owner of the store
	if store.UserID != userID {
		fmt.Println("comparing users", store.UserID, userID)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no such store exists", // more like this user doesn't have such store
		})
		return
	}

	// Parse JSON input and validate it
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// product.ID = utils.NewRandomUUID(c)
	product.StoreID = storeID

	// check if category exists in the store
	// this is done by the gorm Before save function

	// Insert the product into the database
	if err := models.CreateProduct(h.db, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}
