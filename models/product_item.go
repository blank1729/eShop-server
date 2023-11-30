package models

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductItem struct {
	// ProductItemID string    `json:"product_item_id" gorm:"type:uuid;primaryKey;not null"`
	// CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime:nano"`
	// UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
	BaseModel
	Title    *string `json:"title,omitempty"`
	Qunatity uint    `json:"quantity" gorm:"not null;"`
	// SKU
	// IS ARCHIVED
	Price uint `json:"price" gorm:"not null;"`
	// CURRENCY

	ProductID string `json:"product_id" gorm:"not null;type:uuid"`

	OptionVariants []OptionVariant `json:"option_variants" gorm:"many2many:product_item_option_variant;"`
	Images         *[]Image        `json:"images,omitempty" gorm:"foreignKey:ProductItemID;"`

	// ORDERS PLACED
	// isDefault
	// primary_image_id
}

func (productItem *ProductItem) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println("error with the uuid generation", err.Error())
		return err
	}
	productItem.ID = u.String()
	return nil
}

func CreateProductItem(db *gorm.DB, productItem *ProductItem) error {
	err := db.Create(productItem).Error

	if err != nil {
		return err
	}
	err = db.Model(productItem).Association("OptionVariants").Append(productItem.OptionVariants)

	return err
}
