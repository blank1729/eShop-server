package models

type OptionVariation struct {
	BaseModel
	OptionID *string `json:"option_id" gorm:"not null;type:uuid"`

	Value string `json:"value" gorm:"not null;"`

	Products []ProductItem `json:"-" gorm:"many2many:product_item_option_variations;"`
}
