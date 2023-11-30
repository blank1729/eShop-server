package routes

import (
	"eshop/db"
	"eshop/handlers"
	"eshop/middleware"

	"github.com/gin-gonic/gin"
)

func SetProductItemRoutes(router *gin.RouterGroup) {
	productItemHandler := handlers.NewProductItemHandler(db.DB)
	productItemRoutes := router.Group("/stores/:store_id/products/:product_id/productitem")
	{
		// productRoutes.GET("/all", productItemHandler.GetAllProducts)
		// productRoutes.GET("/:product_item_id", productItemHandler.GetProductByID)

		// add IsApiUser Middleware
		productItemRoutes.POST("/", middleware.IsAuthorized(), productItemHandler.CreateProductItem)
		// productRoutes.PUT("/:product_item_id", middleware.IsAuthorized(), productItemHandler.UpdateProduct)
		// productRoutes.DELETE("/:product_item_id", middleware.IsAuthorized(), productItemHandler.DeleteProduct)
	}
}
