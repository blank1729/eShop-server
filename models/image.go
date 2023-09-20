package models

import (
	"gorm.io/gorm"
)

type Image struct {
	// ImageID       uuid.UUID `json:"image_id" gorm:"type:uuid;primaryKey;not null;default:gen_random_uuid()"`
	BaseModel
	Title         string `json:"title"`
	Url           string `json:"url"`
	ProductItemID string `json:"product_id" gorm:"not null;type:uuid"`
}

// CreateImage creates a new image record in the database.
func CreateImage(db *gorm.DB, image *Image) error {
	return db.Create(image).Error
}

// GetImageByID retrieves an image by its ID from the database.
func GetImageByID(db *gorm.DB, imageID string) (*Image, error) {
	var image Image
	err := db.Where("image_id = ?", imageID).First(&image).Error
	if err != nil {
		return nil, err
	}
	return &image, nil
}

// UpdateImage updates an existing image record in the database.
func UpdateImage(db *gorm.DB, image *Image) error {
	return db.Save(image).Error
}

// DeleteImage deletes an image record from the database.
func DeleteImage(db *gorm.DB, image *Image) error {
	return db.Delete(image).Error
}
