package models

import (
	"gorm.io/gorm"
)

type Category struct {
	BaseModel
	Name        string  `json:"name" gorm:"unique" form:"name" binding:"required"`
	Description *string `json:"description,omitempty"`

	StoreID        string  `json:"store_id" gorm:"not null;type:uuid"`
	ParentCategory *string `json:"parent_category,omitempty" gorm:"type:uuid"` // can be null, if not pointer then the null value will be an empty string which is an invalid uuid type

	Products []Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
	Options  []Option  `json:"options,omitempty" gorm:"foreignKey:CategoryID"`
	// image or Icon
}

// CreateCategory creates a new category in the database.
func CreateCategory(db *gorm.DB, category *Category) error {
	return db.Create(category).Error
}

// GetCategoryByID retrieves a category by its ID from the database.
func GetCategoryByID(db *gorm.DB, categoryID string) (*Category, error) {
	var category Category
	if err := db.Where("category_id = ?", categoryID).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func GetAllCategories(db *gorm.DB, store_id string) (*[]Category, error) {
	var categories []Category
	if err := db.Where("store_id = ?", store_id).Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

// UpdateCategory updates an existing category in the database.
func UpdateCategory(db *gorm.DB, existingCategory *Category, updatedCategory *Category) error {
	if err := db.Model(existingCategory).Updates(updatedCategory).Error; err != nil {
		return err
	}
	return nil
}

// DeleteCategory deletes a category by its ID from the database.
func DeleteCategory(db *gorm.DB, categoryID string) error {
	return db.Where("category_id = ?", categoryID).Delete(&Category{}).Error
}

// type Size struct {
// 	gorm.Model
// 	StoreId string `json:"store_id" gorm:"not null;"`
// 	Size    string
// 	Colors  []Color
// }

// type Color struct {
// 	gorm.Model
// 	StoreId string `json:"store_id" gorm:"not null;"`
// 	Color   string
// }

// // Categories []Category
// // Colors     []Color
// // Sizes      []Size
