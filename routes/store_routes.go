package routes

import (
	"eshop/db"
	"eshop/handlers"
	"eshop/middleware"

	"github.com/gin-gonic/gin"
)

func SetStoreRoutes(r *gin.RouterGroup) {
	storesHandler := handlers.NewStoreHandler(db.DB) // Pass your database connection to the handler
	storesRoutes := r.Group("/stores")

	// Define a route to get all stores by UserID
	storesRoutes.POST("/", middleware.IsAuthorized(), storesHandler.CreateStore)

	storesRoutes.GET("/", middleware.IsAuthorized(), storesHandler.GetAllStoresByUserID)
	storesRoutes.GET("/:store_id", storesHandler.GetStore)

	storesRoutes.PUT("/:store_id", storesHandler.UpdateStore)
	storesRoutes.DELETE("/:store_id", storesHandler.DeleteStore)
}
