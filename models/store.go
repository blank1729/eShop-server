package models

import (
	"gorm.io/gorm"
)

type Store struct {
	BaseModel
	Name        string `json:"name" gorm:"size:255;not null" binding:"required" `
	Description string `json:"description" binding:"required"`

	UserID string `json:"-" gorm:"not null;"`

	Categories []Category `json:"categories,omitempty" gorm:"foreignKey:StoreID"`
	Products   []Product  `json:"products,omitempty" gorm:"foreignKey:StoreID"`
	Customers  []Customer `json:"customers,omitempty" gorm:"foreignKey:StoreID"` // Define the foreign key here
	Orders     []Order    `json:"orders,omitempty" gorm:"foreignKey:StoreID"`
	// store Rating
	// store reviews
}

func CreateStore(db *gorm.DB, store *Store) error {
	return db.Create(store).Error
}

// GetStoreByID retrieves a store by its ID from the database.
func GetStoreByID(db *gorm.DB, storeID string) (*Store, error) {
	var store Store
	if err := db.Where("id = ?", storeID).Preload("Categories").Preload("Products").Preload("Customers").Preload("Orders").First(&store).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

// UpdateStore updates an existing store record in the database.
func UpdateStore(db *gorm.DB, store *Store) error {
	return db.Save(store).Error
}

// DeleteStore deletes a store record from the database.
func DeleteStore(db *gorm.DB, store *Store) error {
	return db.Delete(store).Error
}

// FindAllStores retrieves all stores from the database.
// admin route
func FindAllStores(db *gorm.DB) ([]Store, error) {
	// swagger:route GET /api/v1/stores/all pets users listPets
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Deprecated: false
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Parameters:
	//       + name: limit
	//         in: query
	//         description: maximum numnber of results to return
	//         required: false
	//         type: integer
	//         format: int32
	//
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       422: validationError
	var stores []Store
	// .Preload("Categories").Preload("Products").Preload("Customers").Preload("Orders")
	if err := db.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func GetStoresByUserID(db *gorm.DB, userID string) ([]Store, error) {
	var stores []Store
	// .Preload("Categories").Preload("Products").Preload("Customers").Preload("Orders")
	if err := db.Where("user_id = ?", userID).Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}
