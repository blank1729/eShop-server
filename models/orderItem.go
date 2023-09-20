package models

import "gorm.io/gorm"

type OrderItem struct {
	BaseModel
	Quantity      uint    `json:"quantity" form:"quantity" binding:"required,min=1"`
	Price         float64 `json:"items_price" binding:"required,gt=0"`
	OrderID       uint    `json:"order_id" gorm:"not null;type:uuid"`
	ProductItemID string  `json:"product_item_id" gorm:"not null;type:uuid"`
	// Product   Product `gorm:"foreignkey:ProductID"`
	// product variant id
}

// CreateOrderItem creates a new order item record in the database.
func CreateOrderItem(db *gorm.DB, orderItem *OrderItem) error {
	return db.Create(orderItem).Error
}

// GetOrderItemByID retrieves an order item by its ID from the database.
func GetOrderItemByID(db *gorm.DB, orderItemID string) (*OrderItem, error) {
	var orderItem OrderItem
	err := db.Where("order_item_id = ?", orderItemID).First(&orderItem).Error
	if err != nil {
		return nil, err
	}
	return &orderItem, nil
}

// UpdateOrderItem updates an existing order item record in the database.
func UpdateOrderItem(db *gorm.DB, orderItem *OrderItem) error {
	return db.Save(orderItem).Error
}

// DeleteOrderItem deletes an order item record from the database.
func DeleteOrderItem(db *gorm.DB, orderItem *OrderItem) error {
	return db.Delete(orderItem).Error
}
