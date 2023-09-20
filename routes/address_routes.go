package routes

import (
	"eshop/db"
	"eshop/handlers"

	"github.com/gin-gonic/gin"
)

// SetAddressRoutes sets up the routes for addresses.
func SetAddressRoutes(r *gin.RouterGroup) {
	addressRoutes := r.Group("/store/:store_id/users/:user_id/addresses")
	addressHandler := handlers.NewAddressHandler(db.DB)

	// Address routes
	{
		addressRoutes.POST("/", addressHandler.CreateAddress)
		addressRoutes.GET("/:id", addressHandler.GetAddressByID)
		addressRoutes.PUT("/:id", addressHandler.UpdateAddress)
		addressRoutes.DELETE("/:id", addressHandler.DeleteAddress)
	}
}
