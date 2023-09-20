package routes

import (
	"eshop/db"
	"eshop/handlers"
	"eshop/middleware"

	"github.com/gin-gonic/gin"
)

// SetCategoryRoutes defines the routes for the Category resource.
func SetCategoryRoutes(r *gin.RouterGroup) {
	categoryRoutes := r.Group("/stores/:store_id/categories")
	categoryHandler := handlers.NewCategoryHandler(db.DB)
	{
		// add IsApiUser Middleware
		categoryRoutes.POST("/", middleware.IsAuthorized(), categoryHandler.CreateCategory)
		categoryRoutes.PUT("/:category_id", middleware.IsAuthorized(), categoryHandler.UpdateCategory)
		categoryRoutes.DELETE("/:category_id", middleware.IsAuthorized(), categoryHandler.DeleteCategory)

		// get all categories route, public routes
		categoryRoutes.GET("/", categoryHandler.GetAllCategories)
		categoryRoutes.GET("/:category_id", categoryHandler.GetCategoryByID)
	}
}
