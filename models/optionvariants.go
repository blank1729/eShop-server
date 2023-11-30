package models

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OptionVariant struct {
	BaseModel
	OptionID string `json:"option_id" gorm:"not null;type:uuid"`

	Value string `json:"value" gorm:"not null;"`

	Products []ProductItem `json:"-" gorm:"many2many:product_item_option_variant;"`
}

func (OptionVariant *OptionVariant) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println("error with the uuid generation", err.Error())
		return err
	}
	OptionVariant.ID = u.String()
	return nil
}

func CreateOptionVariant(db *gorm.DB, opVariant *OptionVariant) error {
	return db.Create(opVariant).Error
}

func GetAllOptionVariants(db *gorm.DB, optionId string) ([]OptionVariant, error) {
	var optionVariants []OptionVariant
	if err := db.Where("option_id = ?", optionId).Find(&optionVariants).Error; err != nil {
		return nil, err
	}
	return optionVariants, nil
}
