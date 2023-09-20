package handlers

import (
	"eshop/models"
	"eshop/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CustomerHandler is a struct that holds the database connection.
type CustomerHandler struct {
	db *gorm.DB
}

// NewCustomerHandler creates a new CustomerHandler with the provided database connection.
func NewCustomerHandler(db *gorm.DB) *CustomerHandler {
	return &CustomerHandler{db}
}

// CreateCustomer handles the creation of a new customer.
func (h *CustomerHandler) SingUp(c *gin.Context) {

	// Parse JSON input and validate it
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get store id
	store_id := c.Param("store_id")

	// check if the user exists

	// turn password into hashed password
	hashedPassword, _ := utils.GenerateHashPassword(customer.Password)

	u, _ := uuid.NewRandom()

	customer.Password = hashedPassword
	customer.StoreID = store_id
	customer.ID = u.String()

	// Insert the customer into the database
	if err := models.CreateCustomer(h.db, &customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func (h *CustomerHandler) Login(c *gin.Context) {
	// nothing here
	c.JSON(http.StatusOK, gin.H{
		"message": "nothing here",
	})

}

// GetCustomerByID retrieves a customer by their ID.
func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
	customerID := c.Param("id")

	// Retrieve the customer from the database
	customer, err := models.GetCustomerByID(h.db, customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// UpdateCustomer handles the update of an existing customer's information.
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	customerID := c.Param("id")

	// Retrieve the existing customer from the database
	existingCustomer, err := models.GetCustomerByID(h.db, customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	// Parse JSON input and validate it
	var updatedCustomer models.Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the customer's information in the database
	existingCustomer.FirstName = updatedCustomer.FirstName
	existingCustomer.LastName = updatedCustomer.LastName
	existingCustomer.Email = updatedCustomer.Email
	existingCustomer.StoreID = updatedCustomer.StoreID

	if err := models.UpdateCustomer(h.db, existingCustomer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingCustomer)
}

// DeleteCustomer handles the deletion of a customer by their ID.
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	customerID := c.Param("id")

	// Delete the customer from the database
	if err := models.DeleteCustomer(h.db, customerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllCustomers retrieves all customers from the database.
func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
	// Retrieve all customers from the database
	customers, err := models.FindAllCustomers(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}
