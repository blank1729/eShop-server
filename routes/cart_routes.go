package routes

import (
	"eshop/db"
	"eshop/handlers"

	"github.com/gin-gonic/gin"
)

func SetCartRoutes(r *gin.RouterGroup) {
	cartRouter := r.Group("/store/:store_id/users/:user_id/carts")
	cartHandler := handlers.NewCartHandler(db.DB)

	cartRouter.POST("/", cartHandler.CreateCart)
	cartRouter.GET("/:cart_id", cartHandler.GetCartByID)
	cartRouter.PUT("/:cart_id", cartHandler.UpdateCart)
	cartRouter.DELETE("/:cart_id", cartHandler.DeleteCart)
}

func SetCartItemRoutes(r *gin.RouterGroup) {
	cartItemRouter := r.Group("/store/:store_id/users/:user_id/carts/:cart_id/cart_items")
	cartItemHandler := handlers.NewCartItemHandler(db.DB)

	cartItemRouter.POST("/", cartItemHandler.CreateCartItem)
	cartItemRouter.GET("/:cart_item_id", cartItemHandler.GetCartItemByID)
	cartItemRouter.PUT("/:cart_item_id", cartItemHandler.UpdateCartItem)
	cartItemRouter.DELETE("/:cart_item_id", cartItemHandler.DeleteCartItem)
}
