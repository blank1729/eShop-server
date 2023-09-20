package routes

import (
	"eshop/db"
	"eshop/handlers"
	"eshop/middleware"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(r *gin.RouterGroup) {
	userRouter := r.Group("/users")
	userHandler := handlers.NewUserHandler(db.DB) // Pass your database connection to the handler

	// auth routes
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)

	// admin routes === all of the following are admin routes, these are api users not stores users
	userRouter.GET("/all", middleware.IsAuthorized(), userHandler.GetAllUsers)
	// maybe create a checkUser middleware here
	userRouter.GET("/:user_id", middleware.IsAuthorized(), userHandler.GetUserByID)
	userRouter.PUT("/:user_id", middleware.IsAuthorized(), userHandler.UpdateUser)
	userRouter.DELETE("/:user_id", middleware.IsAuthorized(), userHandler.DeleteUser)
}
