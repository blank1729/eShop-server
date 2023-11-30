package handlers

import (
	"gorm.io/gorm"
)

// OptionHandler is a struct that holds the database connection.
type OptionVariantHandler struct {
	db *gorm.DB
}

// NewOptionHandler creates a new OptionHandler with the provided database connection.
func NewOptionVariantHandler(db *gorm.DB) *OptionVariantHandler {
	return &OptionVariantHandler{db}
}
