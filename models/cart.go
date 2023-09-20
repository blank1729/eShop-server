package models

import "gorm.io/gorm"

// each customer can have multiple carts and other people can pay for his cart
type Cart struct {
	BaseModel
	CustomerID uint       `json:"customer_id"`
	CartItems  []CartItem `json:"cart_items" gorm:"foreignkey:CartID"`
	IsPrimary  bool       `json:"is_primary"`
}

// CreateCart creates a new cart record in the database.
func CreateCart(db *gorm.DB, cart *Cart) error {
	return db.Create(cart).Error
}

// GetCartByID retrieves a cart by its ID from the database.
func GetCartByID(db *gorm.DB, cartID string) (*Cart, error) {
	var cart Cart
	err := db.Where("id = ?", cartID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

// UpdateCart updates an existing cart record in the database.
func UpdateCart(db *gorm.DB, cart *Cart) error {
	return db.Save(cart).Error
}

// DeleteCart deletes a cart record from the database.
func DeleteCart(db *gorm.DB, cart *Cart) error {
	return db.Delete(cart).Error
}

// ==========================

type CartItem struct {
	BaseModel
	Quantity      uint   `json:"quantity" form:"quantity" binding:"required,min=1"`
	ProductItemID string `json:"product_item_id" gorm:"not null;type:uuid"`
	CartID        string `json:"cart_id" gorm:"not null;type:uuid"`
}

// CreateCartItem creates a new cart item record in the database.
func CreateCartItem(db *gorm.DB, cartItem *CartItem) error {
	return db.Create(cartItem).Error
}

// GetCartItemByID retrieves a cart item by its ID from the database.
func GetCartItemByID(db *gorm.DB, cartItemID string) (*CartItem, error) {
	var cartItem CartItem
	err := db.Where("id = ?", cartItemID).First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

// UpdateCartItem updates an existing cart item record in the database.
func UpdateCartItem(db *gorm.DB, cartItem *CartItem) error {
	return db.Save(cartItem).Error
}

// DeleteCartItem deletes a cart item record from the database.
func DeleteCartItem(db *gorm.DB, cartItem *CartItem) error {
	return db.Delete(cartItem).Error
}
