package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// create a new random uuid and incase of an error returns status code 500
func NewRandomUUID(c *gin.Context) string {
	u, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create uuid",
		})
	}
	return u.String()
}
