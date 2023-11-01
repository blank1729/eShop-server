package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckUser(c *gin.Context, user_id string) bool {
	u, exists := c.Get("user_id")
	if !exists {
		return false
	}
	fmt.Println("fromt context -> ", u, "from uri -> ", user_id)
	return u == user_id
	// can write it this wasy
	// return u == user_id && exists
}
