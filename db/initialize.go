package db

import (
	"eshop/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	dsn := "host=localhost user=postgres password=password dbname=eshop_test sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("cannot connect to db")
		log.Fatal(err)
	}
}

func SyncDatabase() {
	// Assuming you have a GORM DB instance named "db"

	// AutoMigrate all the models
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Store{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Option{})
	DB.AutoMigrate(&models.OptionVariant{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.ProductItem{})
	DB.AutoMigrate(&models.Image{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.CartItem{})
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.OrderItem{})
}
