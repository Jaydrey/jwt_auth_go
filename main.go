package main

import (
	"jwt_auth_go/controllers"
	"jwt_auth_go/initializers"
	"jwt_auth_go/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "API Working"})
}

func main() {
	router := gin.Default()

	router.GET("/ping", ping)
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run()
}
