package models

import (
	"gorm.io/gorm"
)

type Order struct {
	BaseModel
	TotalPrice    float64     `json:"total_price" binding:"required,gt=0"`
	OrderStatus   string      `json:"order_status" gorm:"default:'pending"`
	PaymentStatus string      `json:"payment_status" gorm:"default:'pending"`
	StoreID       string      `json:"store_id" gorm:"not null;type:uuid"`
	CustomerID    string      `json:"customer_id" gorm:"not null;type:uuid"`
	AddressID     string      `json:"address_id" gorm:"not null;type:uuid"`
	Items         []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	// ordered by another user
	// payment method
	// date and time of payment
	// shipped datetime
	// delivery date time
}

// CreateOrder creates a new order record in the database.
func CreateOrder(db *gorm.DB, order *Order) error {
	return db.Create(order).Error
}

// GetOrderByID retrieves an order by its ID from the database.
func GetOrderByID(db *gorm.DB, orderID string) (*Order, error) {
	var order Order
	err := db.Where("id = ?", orderID).Preload("Items").First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateOrder updates an existing order record in the database.
func UpdateOrder(db *gorm.DB, order *Order) error {
	return db.Save(order).Error
}

// DeleteOrder deletes an order record from the database.
func DeleteOrder(db *gorm.DB, order *Order) error {
	return db.Delete(order).Error
}
