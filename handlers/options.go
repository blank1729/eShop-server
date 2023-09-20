package handlers

import (
	"net/http"

	"eshop/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OptionHandler is a struct that holds the database connection.
type OptionHandler struct {
	db *gorm.DB
}

// NewOptionHandler creates a new OptionHandler with the provided database connection.
func NewOptionHandler(db *gorm.DB) *OptionHandler {
	return &OptionHandler{db}
}

// CreateOption handles the creation of a new option.
func (h *OptionHandler) CreateOption(c *gin.Context) {
	// Extract store ID from the URL
	storeID := c.Param("store_id")

	// Parse JSON input and validate it
	var option models.Option
	if err := c.ShouldBindJSON(&option); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// extract user_id from token

	// check if the user has access

	// check the user access controls

	// check if the option exists
	var existingOption models.Option
	if err := h.db.Where("name = ?", option.Name).
		Where("category_id = ?", option.CategoryID).
		First(&existingOption).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Option exists for this category"})
		return
	}

	u, _ := uuid.NewRandom()
	option.ID = u.String()

	// Retrieve the category from the database based on the option's CategoryID
	category, err := models.GetCategoryByID(h.db, option.CategoryID)
	if err != nil {
		// since category does not exists we have check if the store exists

		if _, err := models.GetCategoryByID(h.db, storeID); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Store not found"})
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Check if the store ID from the URL matches the store ID in the category
	if category.StoreID != storeID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Category does not belong to the store"})
		return
	}

	if option.OptionVariations != nil {
		for _, op := range option.OptionVariations {
			u, _ = uuid.NewRandom()
			x := u.String()
			op.ID = x
		}
	}

	// Insert the option into the database
	if err := models.CreateOption(h.db, &option); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, option)
}

// GetOptionByID retrieves an option by its ID.
func (h *OptionHandler) GetOptionByID(c *gin.Context) {
	// Extract store ID from the URL
	storeID := c.Param("store_id")

	optionID := c.Param("id")

	// Retrieve the option from the database
	option, err := models.GetOptionByID(h.db, optionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
		return
	}

	// Retrieve the category from the database based on the option's CategoryID
	category, err := models.GetCategoryByID(h.db, option.CategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Check if the store ID from the URL matches the store ID in the category
	if category.StoreID != storeID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Category does not belong to the store"})
		return
	}

	c.JSON(http.StatusOK, option)
}

// UpdateOption updates an existing option.
func (h *OptionHandler) UpdateOption(c *gin.Context) {
	// Extract store ID from the URL
	storeID := c.Param("store_id")

	optionID := c.Param("id")

	// Retrieve the option from the database
	existingOption, err := models.GetOptionByID(h.db, optionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
		return
	}

	// Retrieve the category from the database based on the option's CategoryID
	category, err := models.GetCategoryByID(h.db, existingOption.CategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Check if the store ID from the URL matches the store ID in the category
	if category.StoreID != storeID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Category does not belong to the store"})
		return
	}

	// Parse JSON input and validate it
	var updatedOption models.Option
	if err := c.ShouldBindJSON(&updatedOption); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the option in the database
	if err := models.UpdateOption(h.db, existingOption, &updatedOption); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingOption)
}

// DeleteOption deletes an option by its ID.
func (h *OptionHandler) DeleteOption(c *gin.Context) {
	// Extract store ID from the URL
	storeID := c.Param("store_id")

	optionID := c.Param("id")

	// Retrieve the option from the database
	existingOption, err := models.GetOptionByID(h.db, optionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
		return
	}

	// Retrieve the category from the database based on the option's CategoryID
	category, err := models.GetCategoryByID(h.db, existingOption.CategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Check if the store ID from the URL matches the store ID in the category
	if category.StoreID != storeID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Category does not belong to the store"})
		return
	}

	// Delete the option from the database
	if err := models.DeleteOption(h.db, optionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
