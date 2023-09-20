package handlers

import (
	"errors"
	"eshop/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProductHandler is a struct that holds the database connection.
type ProductHandler struct {
	db *gorm.DB
}

// NewProductHandler creates a new ProductHandler with the provided database connection.
func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{db}
}

// CreateProduct handles the creation of a new product.
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	storeID := c.Param("store_id")
	userID, _ := c.Get("user_id") // not using exists here since if user_id doesn't exists then it shouldn't pass the middleware

	// check if the store exists
	store, err := models.GetStoreByID(h.db, storeID)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Store doesn't exist",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	// check if the user is the owner of the store
	if store.UserID != userID {
		fmt.Println("comparing users", store.UserID, userID)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no such store exists", // more like this user doesn't have such store
		})
		return
	}

	// Parse JSON input and validate it
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// product.ID = utils.NewRandomUUID(c)
	product.StoreID = storeID

	// check if category exists in the store
	// this is done by the gorm Before save function

	// Insert the product into the database
	if err := models.CreateProduct(h.db, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProductByID retrieves a product by its ID.
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productID := c.Param("id")

	// Retrieve the product from the database
	product, err := models.GetProductByID(h.db, productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct handles the update of an existing product's information.
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	// Retrieve the existing product from the database
	existingProduct, err := models.GetProductByID(h.db, productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// var p models.Product
	// json.NewDecoder(c.Request.Body).Decode(&p)
	// c.JSON(http.StatusOK, p)

	// fmt.Println(existingProduct)

	// Parse JSON input and validate it
	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the product's information in the database
	existingProduct.Name = updatedProduct.Name
	existingProduct.Description = updatedProduct.Description
	existingProduct.Rating = updatedProduct.Rating

	store_id := c.Param("store_id")
	updatedProduct.StoreID = store_id
	for _, p := range updatedProduct.ProductItems {
		fmt.Println("product item id ", p.ID)
		u, _ := uuid.NewRandom()
		x := u.String()
		p.ID = x
		p.ProductID = productID
		// fmt.Println(h.db.Create(p).Error)
		fmt.Printf("%+v\n", p)

		for _, ov := range p.OptionVariations {
			u, _ := uuid.NewRandom()
			x := u.String()
			ov.ID = x
		}

		existingProduct.ProductItems = append(existingProduct.ProductItems, p)
	}

	// existingProduct.ProductItems = updatedProduct.ProductItems

	if err := models.UpdateProduct(h.db, existingProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingProduct)
}

// DeleteProduct handles the deletion of a product by its ID.
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	// Delete the product from the database
	if err := models.DeleteProduct(h.db, productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllProducts retrieves all products from the database for a specific store.
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	// Extract store ID from the URL
	storeID := c.Param("store_id")

	// Retrieve all products for the specified store
	products, err := models.FindAllProducts(h.db, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// func Custom() {
// 	// Check if the category already exists in the database
// 	var existingCategory models.Category
// 	if err := db.DB.First(&existingCategory, "name = ?", "mens").Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			// Category doesn't exist, create a new one
// 			c1 := models.Category{
// 				Name: "mens",
// 			}
// 			if result := db.DB.Create(&c1); result.Error != nil {
// 				fmt.Println("Error creating category:", result.Error)
// 				return
// 			}
// 			fmt.Println("Created new category:", c1.Name)
// 			// Now you can proceed to create the product with the new category
// 		} else {
// 			fmt.Println("Error querying category:", err)
// 			return
// 		}
// 	}

// 	// If the category already exists or has been created, create the product
// 	p := models.Product{
// 		Name:        "Pants",
// 		Description: "",
// 		// Categories:  []models.Category{existingCategory},
// 	}

// 	result := db.DB.Create(&p)
// 	if result.Error != nil {
// 		fmt.Println("Error creating product:", result.Error)
// 		return
// 	}

// 	fmt.Println("Created db entry for product:", p.Name)

// }
