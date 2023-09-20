package handlers

import (
	"errors"
	"eshop/models"
	"eshop/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	db *gorm.DB
}

// NewCategoryHandler creates a new CategoryHandler with the provided database connection.
func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{db}
}

// CreateCategory handles the creation of a new category.
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	// Parse JSON input and validate it
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeID := c.Param("store_id")

	// not using exists here since if user_id doesn't exists then it shouldn't pass the middleware
	userID, _ := c.Get("user_id")

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

	// user and store exists, he owns the store
	category.ID = utils.NewRandomUUID(c)
	category.StoreID = storeID

	// extract store id
	// store_id := category.StoreID
	var store_id = c.Param("store_id")
	category.StoreID = store_id

	category.ID = utils.NewRandomUUID(c)

	// Insert the category into the database
	if err := models.CreateCategory(h.db, &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   nil,
		"message": "created category",
		"data":    category,
	})
}

// GetCategoryByID retrieves a category by its ID.
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	categoryID := c.Param("id")

	// Retrieve the category from the database
	category, err := models.GetCategoryByID(h.db, categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates an existing category.
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")

	// Retrieve the category from the database
	existingCategory, err := models.GetCategoryByID(h.db, categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Parse JSON input and validate it
	var updatedCategory models.Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the category in the database
	if err := models.UpdateCategory(h.db, existingCategory, &updatedCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingCategory)
}

// DeleteCategory deletes a category by its ID.
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	// Delete the category from the database
	if err := models.DeleteCategory(h.db, categoryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteCategory deletes a category by its ID.
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {

	store_id := c.Param("store_id")

	// Delete the category from the database
	categories, err := models.GetAllCategories(h.db, store_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
