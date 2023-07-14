package main

import (
	"github.com/gin-gonic/gin"
	"jwt-api/controllers"
	"jwt-api/middlewares"
	"jwt-api/models"
)

func main() {
	models.ConnectDB()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
