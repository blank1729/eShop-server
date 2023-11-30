package models

import (
	"gorm.io/gorm"
)

type Option struct {
	BaseModel
	Name        string  `json:"name" gorm:"not null,unique"`
	Description *string `json:"description"`

	CategoryID string `json:"category_id" gorm:"not null;type:uuid"`

	OptionVariants *[]OptionVariant `json:"option_variants" gorm:"foreignKey:OptionID"`
}

// CreateOption creates a new option in the database.
func CreateOption(db *gorm.DB, option *Option) error {
	return db.Create(option).Error
}

// GetOptionByID retrieves an option by its ID from the database.
func GetOptionByID(db *gorm.DB, optionID string) (*Option, error) {
	var option Option
	if err := db.Where("id = ?", optionID).Preload("OptionVariants").First(&option).Error; err != nil {
		return nil, err
	}
	return &option, nil
}

// UpdateOption updates an existing option in the database.
func UpdateOption(db *gorm.DB, existingOption *Option, updatedOption *Option) error {
	if err := db.Model(existingOption).Updates(updatedOption).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOption deletes an option by its ID from the database.
func DeleteOption(db *gorm.DB, optionID string) error {
	return db.Where("id = ?", optionID).Delete(&Option{}).Error
}
