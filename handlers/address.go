package handlers

import (
	"eshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddressHandler handles HTTP requests for the Address model.
type AddressHandler struct {
	db *gorm.DB
}

// NewAddressHandler creates a new AddressHandler with the provided database connection.
func NewAddressHandler(db *gorm.DB) *AddressHandler {
	return &AddressHandler{db}
}

// CreateAddress handles the creation of a new address.
func (h *AddressHandler) CreateAddress(c *gin.Context) {
	var address models.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateAddress(h.db, &address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, address)
}

// GetAddressByID retrieves an address by its ID.
func (h *AddressHandler) GetAddressByID(c *gin.Context) {
	addressID := c.Param("id")
	address, err := models.GetAddressByID(h.db, addressID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}
	c.JSON(http.StatusOK, address)
}

// UpdateAddress updates an existing address.
func (h *AddressHandler) UpdateAddress(c *gin.Context) {
	addressID := c.Param("id")
	var updatedAddress models.Address
	if err := c.ShouldBindJSON(&updatedAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	address, err := models.GetAddressByID(h.db, addressID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	// Update address fields here
	address.FirstName = updatedAddress.FirstName
	address.LastName = updatedAddress.LastName
	address.House = updatedAddress.House
	address.Street = updatedAddress.Street
	address.City = updatedAddress.City
	address.State = updatedAddress.State
	address.Country = updatedAddress.Country
	address.PostalCode = updatedAddress.PostalCode
	address.IsDefault = updatedAddress.IsDefault

	if err := models.UpdateAddress(h.db, address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, address)
}

// DeleteAddress deletes an address.
func (h *AddressHandler) DeleteAddress(c *gin.Context) {
	addressID := c.Param("id")
	address, err := models.GetAddressByID(h.db, addressID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address not found"})
		return
	}

	if err := models.DeleteAddress(h.db, address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
