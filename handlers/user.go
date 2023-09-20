package handlers

import (
	"eshop/models"
	"eshop/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func ValidateRole(role string) bool {
// 	return role == "user" || role == "admin"
// }

// func ValidateRoleField(fl validator.FieldLevel) bool {
// 	role := fl.Field().String()
// 	return ValidateRole(role)
// }

// validate := validator.New()

// update user
type UpdateUserRequest struct {
	Username *string `json:"username" binding:"omitempty,min=2"`
	Password *string `json:"password" binding:"omitempty,min=10"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Role     *string `json:"role" binding:"omitempty"` // write custom validator
}

type UserHandler struct {
	db *gorm.DB // Your database connection
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	if !(utils.CheckUser(c, userID)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unautorized"})
		return
	}

	user, err := models.GetUserByID(h.db, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var updateReq UpdateUserRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var user models.User
	// Check if the user exists
	user, err := models.GetUserByID(h.db, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	fmt.Println(user)

	// Update the user
	if updateReq.Username != nil {
		user.Username = *updateReq.Username
	}
	if updateReq.Password != nil {
		user.Password, _ = utils.GenerateHashPassword(*updateReq.Password)
	}
	if updateReq.Email != nil {
		user.Email = *updateReq.Email
	}
	// check for admin status
	// if updateReq.Role != nil {
	// 	user.Role = *updateReq.Role
	// }

	h.db.Save(&user)

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := models.GetUserByID(h.db, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := models.DeleteUser(h.db, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := h.db.Preload("Stores").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
