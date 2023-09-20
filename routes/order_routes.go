package routes

import (
	"eshop/db"
	"eshop/handlers"

	"github.com/gin-gonic/gin"
)

// SetOrderRoutes sets up the routes for orders.
func SetOrderRoutes(r *gin.RouterGroup) {
	orderRoutes := r.Group("/store/:store_id/users/:user_id/orders")
	orderHandler := handlers.NewOrderHandler(db.DB)

	// Orders routes
	{
		orderRoutes.POST("/", orderHandler.CreateOrder)
		orderRoutes.GET("/:id", orderHandler.GetOrderByID)
		orderRoutes.PUT("/:id", orderHandler.UpdateOrder)
		orderRoutes.DELETE("/:id", orderHandler.DeleteOrder)
	}
}
