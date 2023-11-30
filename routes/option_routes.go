package routes

import (
	"eshop/db"
	"eshop/handlers"

	"github.com/gin-gonic/gin"
)

func SetOptionRoutes(r *gin.RouterGroup) {
	optionsRouter := r.Group("/stores/:store_id/options")

	optionHandler := handlers.NewOptionHandler(db.DB)

	optionsRouter.POST("/", optionHandler.CreateOption)
	optionsRouter.GET("/:option_id", optionHandler.GetOptionByID)
	optionsRouter.PUT("/:option_id", optionHandler.UpdateOption)
	optionsRouter.DELETE("/:option_id", optionHandler.DeleteOption)
}
