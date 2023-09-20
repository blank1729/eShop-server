package handlers

import (
	"net/http"

	"eshop/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ImageHandler handles HTTP requests for the Image model.
type ImageHandler struct {
	db *gorm.DB
}

// NewImageHandler creates a new ImageHandler with the provided database connection.
func NewImageHandler(db *gorm.DB) *ImageHandler {
	return &ImageHandler{db}
}

// CreateImage handles the creation of a new image.
func (h *ImageHandler) CreateImage(c *gin.Context) {
	var image models.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateImage(h.db, &image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, image)
}

// GetImageByID retrieves an image by its ID.
func (h *ImageHandler) GetImageByID(c *gin.Context) {
	imageID := c.Param("id")
	image, err := models.GetImageByID(h.db, imageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.JSON(http.StatusOK, image)
}

// UpdateImage updates an existing image.
func (h *ImageHandler) UpdateImage(c *gin.Context) {
	imageID := c.Param("id")
	var updateImage models.Image

	if err := c.ShouldBindJSON(&updateImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the image exists
	existingImage, err := models.GetImageByID(h.db, imageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// Update the image attributes
	existingImage.Title = updateImage.Title
	existingImage.Url = updateImage.Url

	if err := models.UpdateImage(h.db, existingImage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingImage)
}

// DeleteImage deletes an image by its ID.
func (h *ImageHandler) DeleteImage(c *gin.Context) {
	imageID := c.Param("id")

	// Check if the image exists
	image, err := models.GetImageByID(h.db, imageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// Delete the image
	if err := models.DeleteImage(h.db, image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
