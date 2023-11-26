package models

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	BaseModel
	// ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;not null"`

	Name        string  `json:"name" binding:"required,min=1"`
	Description string  `json:"description" binding:"required,min=1"`
	Rating      float32 `json:"rating" binding:"gte=0,lte=5" gorm:"default:0"`

	CategoryID string  `json:"category_id" gorm:"not null;type:uuid" binding:"required"`
	StoreID    string  `json:"store_id" gorm:"not null;type:uuid"`
	ImageID    *string `json:"image_id,omitempty"`

	ProductItems []ProductItem `json:"product_items,omitempty" gorm:"foreignKey:ProductID"`
	// Reviews
	// Ratings by users
	// overall rating -> Rating
	// is archived
	// is featured
	// is active
	// primary_image_id
	// brand_id
	// Code
	// price *int = can be null
	// DiscountPrice
	// weight
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println("error with the uuid generation", err.Error())
		return err
	}
	product.ID = u.String()
	return nil
}

func (product *Product) BeforeSave(tx *gorm.DB) (err error) {
	// Check if the StoreID in the Product matches the StoreID in the associated Category
	var category Category
	if tx.Where("id = ?", product.CategoryID).First(&category).Error == gorm.ErrRecordNotFound {
		return errors.New("Category not found")
	}

	if category.StoreID != product.StoreID {
		return errors.New("StoreID in Product does not match StoreID in Category")
	}

	return nil
}

// type ProductItemOptionVariation struct {
// 	ProductItemID     string `json:"product_item_id" gorm:"type:uuid;primaryKey;not null"`
// 	OptionVariationID string `json:"option_variation_id" gorm:"type:uuid;primaryKey;not null"`
// }

// CreateProduct creates a new product record in the database.
func CreateProduct(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}

// GetProductByID retrieves a product by its ID.
func GetProductByID(db *gorm.DB, productID string) (*Product, error) {
	var product Product
	if err := db.Where("product_id = ?", productID).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product's information in the database.
func UpdateProduct(db *gorm.DB, product *Product) error {
	return db.Save(product).Error
}

// DeleteProduct deletes a product by its ID.
func DeleteProduct(db *gorm.DB, productID string) error {
	return db.Where("product_id = ?", productID).Delete(&Product{}).Error
}

// FindAllProducts retrieves all products from the database.
func FindAllProducts(db *gorm.DB, storeID string) ([]Product, error) {
	var products []Product
	if err := db.Where("store_id = ?", storeID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

type ProductWithItems struct {
	Product Product
	Items   []ProductItem
}

func GetProductsByQuery(db *gorm.DB, queryParams map[string]string) ([]ProductWithItems, error) {
	var results []ProductWithItems

	// Fetch products based on query parameters.
	err := db.Preload("ProductItems", func(db *gorm.DB) *gorm.DB {
		// Apply filters to ProductItems based on query parameters.
		for key, value := range queryParams {
			switch key {
			case "size":
				db = db.Where("product_items.size = ?", value)
			case "color":
				db = db.Where("product_items.color = ?", value)
			case "price":
				db = db.Where("product_items.price <= ?", value)
				// Add more cases for additional filters as needed.
			}
		}
		return db
	}).Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
