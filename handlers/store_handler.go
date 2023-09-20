package handlers

import (
	"eshop/db"
	"eshop/models"
	"eshop/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreHandler struct {
	db *gorm.DB // Your database connection
}

func NewStoreHandler(db *gorm.DB) *StoreHandler {
	return &StoreHandler{db}
}

// CreateStore creates a new store
func (h *StoreHandler) CreateStore(c *gin.Context) {
	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// extract user id
	user_id, _ := c.Get("user_id")

	// check if the mentioned in the request is the actual user
	// if !utils.CheckUser(c, user_id) {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	// 	return
	// }

	// check if the user exists
	exists, err := models.UserExists(db.DB, user_id.(string))
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user does not exist"})
		return
	}
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	// check the user access controls like the number of stores he can create and such

	store.ID = utils.NewRandomUUID(c)
	fmt.Println(store.ID)
	store.UserID = user_id.(string)

	if err := db.DB.Create(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   nil,
		"message": "created store",
		"data":    store,
	})
}

// GetStore retrieves a store by ID
func (h *StoreHandler) GetStore(c *gin.Context) {
	storeID := c.Param("id")

	var store models.Store
	if err := db.DB.Where("store_id = ?", storeID).Preload("Categories").Preload("Products").Preload("Customers").Preload("Orders").First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	c.JSON(http.StatusOK, store)
}

// UpdateStore updates an existing store by ID
func (h *StoreHandler) UpdateStore(c *gin.Context) {
	storeID := c.Param("id")

	var store models.Store
	if err := db.DB.Where("store_id = ?", storeID).First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, store)
}

// DeleteStore deletes a store by ID
func (h *StoreHandler) DeleteStore(c *gin.Context) {
	storeID := c.Param("id")

	var store models.Store
	if err := db.DB.Where("store_id = ?", storeID).First(&store).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	if err := db.DB.Delete(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetAllStoresByUserID retrieves all stores by UserID
func (h *StoreHandler) GetAllStoresByUserID(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var stores []models.Store
	if err := db.DB.Where("user_id = ?", userID.(string)).Find(&stores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": nil,
		"err":     nil,
		"data":    stores,
	})
}
