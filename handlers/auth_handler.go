package handlers

import (
	"eshop/db"
	"eshop/models"
	"eshop/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=10"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// unmarshsal the signup request from client
	var request SignupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if there an existing user with give email or username
	if _, err := models.GetUserByEmail(h.db, request.Email); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user exists"})
		return
	}

	// create a user object
	var user models.User

	u, _ := uuid.NewRandom()

	user.ID = u.String()
	user.Email = request.Email
	user.Username = request.Username
	user.Password, _ = utils.GenerateHashPassword(request.Password)

	// Add user to database
	if err := models.CreateUser(h.db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"error":   nil,
		"user_id": user.ID,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	user := LoginRequest{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid format",
		})
		return
	}

	// check if the user exists
	existingUser, err := models.GetUserByEmail(h.db, user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user does not exist",
		})
		return
	}

	// if the user does exists then check the passowrd
	if !(utils.CompareHashPassword(user.Password, existingUser.Password)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Wrong username or password",
		})
		return
	}

	// JWT stuff
	expirationTime := time.Now().Add(time.Hour * 24 * 7)

	claims := &models.Claims{
		Role:   existingUser.Role,
		UserID: existingUser.ID,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(utils.JWTkey)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	// cookie stuff
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"message": "logged in",
		"error":   nil,
		"data": gin.H{
			"user_id": existingUser.ID,
		},
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}

func CreateUser(c *gin.Context) {
	user := SignupRequest{}

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// check if the user exists
	existingUser := models.User{}

	db.DB.Where("email", user.Email).First(&existingUser)
	if existingUser.ID != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User exists",
		})
		return
	}

	db.DB.Where("username", user.Username).First(&existingUser)
	if existingUser.ID != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "usename already exists",
		})
		return
	}

	// add user to database
	u, _ := uuid.NewRandom()
	newUser := models.User{}

	newUser.ID = u.String()
	newUser.Username = user.Username
	newUser.Email = user.Email
	newUser.Password = user.Password

	db.DB.Create(&newUser)

	// send 201 created
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
	})

}
