package main

import (
	"eshop/db"
	"eshop/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	db.Initialize()
	db.SyncDatabase()
}
func main() {

	// gin server

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://example.com", "http://localhost:3000"} // List of allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}     // Allowed HTTP methods
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}     // Allowed headers
	r.Use(cors.New(config))

	apiRouter := r.Group("/api/v1")
	routes.Routes(apiRouter)

	r.Run(":8080")

}
