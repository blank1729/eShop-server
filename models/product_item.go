package models

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

	OptionVariations []OptionVariation `json:"product_item_option_variations" gorm:"many2many:product_item_option_variations;"`
	Images           []Image           `json:"images,omitempty" gorm:"foreignKey:ProductItemID;"`

	// ORDERS PLACED
	// isDefault
	// primary_image_id
}
