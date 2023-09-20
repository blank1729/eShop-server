package routes

import (
	"eshop/db"
	"eshop/handlers"
	"eshop/middleware"

	"github.com/gin-gonic/gin"
)

func SetProductRoutes(router *gin.RouterGroup) {
	productHandler := handlers.NewProductHandler(db.DB)
	productRoutes := router.Group("/stores/:store_id/products")
	{
		productRoutes.GET("/all", productHandler.GetAllProducts)
		productRoutes.GET("/:product_id", productHandler.GetProductByID)

		// add IsApiUser Middleware
		productRoutes.POST("/", middleware.IsAuthorized(), productHandler.CreateProduct)
		productRoutes.PUT("/:product_id", middleware.IsAuthorized(), productHandler.UpdateProduct)
		productRoutes.DELETE("/:product_id", middleware.IsAuthorized(), productHandler.DeleteProduct)
	}
}
