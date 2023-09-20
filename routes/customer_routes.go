package routes

import (
	"eshop/db"
	"eshop/handlers"

	"github.com/gin-gonic/gin"
)

func SetCustomerRoutes(router *gin.RouterGroup) {
	customerHandler := handlers.NewCustomerHandler(db.DB)
	customerRoutes := router.Group("/stores/:store_id/")
	{
		// auth web hook for each store
		// customerRoutes.POST("/auth/github", customerHandler.AuthHandler)
		// customerRoutes.POST("/auth/google", customerHandler.AuthHandler)
		// customerRoutes.POST("/auth/", customerHandler.AuthHandler)

		customerRoutes.POST("/signup", customerHandler.SingUp)
		customerRoutes.POST("/login", customerHandler.Login)

		// middleware IsStoreUser
		customerRoutes.GET("/:user_id", customerHandler.GetCustomerByID)
		customerRoutes.PUT("/:user_id", customerHandler.UpdateCustomer)
		customerRoutes.DELETE("/:user_id", customerHandler.DeleteCustomer)

		// add IsApiUser Middleware
		customerRoutes.GET("/", customerHandler.GetAllCustomers)
	}
}
