package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	BaseModel
	FirstName string    `json:"first_name" gorm:"size:255;not null" binding:"required,min=2" `
	LastName  string    `json:"last_name" gorm:"size:255;not null" binding:"required,min=2"`
	Email     string    `json:"email" gorm:"index;not null;unique" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=10"`
	StoreID   string    `json:"store_id" gorm:"type:uuid;not null"`
	Addresses []Address `json:"addresses,omitempty" gorm:"foreignKey:CustomerID"`
	Orders    []Order   `json:"orders,omitempty" gorm:"foreignKey:CustomerID"`
	Carts     []Cart    `json:"carts,omitempty" gorm:"foreignKey:CustomerID"`
	// reviews
	// ratings
	// Searches
	// product visits
}

// CreateCustomer creates a new customer record in the database.
func CreateCustomer(db *gorm.DB, customer *Customer) error {
	return db.Create(customer).Error
}

// GetCustomerByID retrieves a customer by their ID.
func GetCustomerByID(db *gorm.DB, customerID string) (*Customer, error) {
	var customer Customer
	if err := db.Where("id = ?", customerID).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

// UpdateCustomer updates an existing customer's information in the database.
func UpdateCustomer(db *gorm.DB, customer *Customer) error {
	return db.Save(customer).Error
}

// DeleteCustomer deletes a customer by their ID.
func DeleteCustomer(db *gorm.DB, customerID string) error {
	return db.Where("id = ?", customerID).Delete(&Customer{}).Error
}

// FindCustomerByEmail finds a customer by their email.
func FindCustomerByEmail(db *gorm.DB, email string) (*Customer, error) {
	var customer Customer
	if err := db.Where("email = ?", email).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

// FindAllCustomers retrieves all customers from the database.
func FindAllCustomers(db *gorm.DB) ([]Customer, error) {
	var customers []Customer
	if err := db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
