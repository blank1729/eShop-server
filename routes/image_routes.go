package routes

import (
	"eshop/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetImageRoutes sets up the image routes.
func SetImageRoutes(r *gin.Engine, db *gorm.DB) {
	imageRoutes := r.Group("/stores/:store_id/products/:product_id/images")
	imageHandler := handlers.NewImageHandler(db)

	imageRoutes.POST("/", imageHandler.CreateImage)
	imageRoutes.GET("/:id", imageHandler.GetImageByID)
	imageRoutes.PUT("/:id", imageHandler.UpdateImage)
	imageRoutes.DELETE("/:id", imageHandler.DeleteImage)
}
