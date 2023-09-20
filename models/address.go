package models

import "gorm.io/gorm"

type Address struct {
	BaseModel
	FirstName  string `json:"first_name" binding:"required" gorm:"size:255;not null"`
	LastName   string `json:"last_name" binding:"required" gorm:"size:255;not null"`
	House      string `json:"house_no" form:"house_no" binding:"required"`
	Street     string `json:"street" form:"street" binding:"required"`
	City       string `json:"city" form:"city" binding:"required"`
	State      string `json:"state" form:"state" binding:"required"`
	Country    string `json:"country" form:"country" binding:"required"`
	PostalCode string `json:"postal_code" form:"postal_code" binding:"required"`
	IsDefault  bool   `json:"is_default"`
	CustomerID string `json:"customer_id" gorm:"not null;type:uuid"`
}

// CreateAddress creates a new address record in the database.
func CreateAddress(db *gorm.DB, address *Address) error {
	return db.Create(address).Error
}

// GetAddressByID retrieves an address by its ID from the database.
func GetAddressByID(db *gorm.DB, addressID string) (*Address, error) {
	var address Address
	err := db.Where("id = ?", addressID).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

// UpdateAddress updates an existing address record in the database.
func UpdateAddress(db *gorm.DB, address *Address) error {
	return db.Save(address).Error
}

// DeleteAddress deletes an address record from the database.
func DeleteAddress(db *gorm.DB, address *Address) error {
	return db.Delete(address).Error
}
