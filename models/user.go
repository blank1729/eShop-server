package models

import (
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	FirstName string  `json:"first_name,omitempty" gorm:"size:255;not null" binding:"required" `
	LastName  string  `json:"last_name,omitempty" gorm:"size:255;not null" binding:"required" `
	Username  string  `json:"username" gorm:"index;not null;unique"`
	Email     string  `json:"email" gorm:"index;not null;unique"`
	Password  string  `json:"-"`
	Role      string  `json:"role" gorm:"index;not null;default:'user'"`
	Stores    []Store `json:"stores,omitempty" gorm:"foreignKey:UserID" binding:"omitempty"`
	// payment details
	// store count
	// products count
	// current plan
	// on trial
	// trial data start and end
}

// CreateUser creates a new user record in the database.
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// GetUserByID retrieves a user by ID from the database.
func GetUserByID(db *gorm.DB, userID string) (*User, error) {
	var user User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by email from the database.
func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user record in the database.
func UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error
}

// DeleteUser deletes a user record from the database.
func DeleteUser(db *gorm.DB, user *User) error {
	return db.Delete(user).Error
}

func UserExists(db *gorm.DB, user_id string) (bool, error) {
	var user User
	err := db.Where("id = ?", user_id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		// User does not exist
		return false, nil
	} else if err != nil {
		// user still doesn't exists but we don't have do check it
		return false, err
	} else {
		return true, nil
	}

}
