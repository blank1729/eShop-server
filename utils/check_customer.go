package utils

import "github.com/gin-gonic/gin"

func CheckCustomer(c *gin.Context, customer_id string) bool {
	cid, exists := c.Get("customer_id")
	if !exists {
		return false
	}
	return cid == customer_id
}
