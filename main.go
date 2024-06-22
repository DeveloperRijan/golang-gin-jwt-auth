package main

import (
	"fmt"
	"golang_gin_jwt_auth/handlers"
	"golang_gin_jwt_auth/initializers"
	"golang_gin_jwt_auth/models"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnect()
	//migrate tables
	initializers.DB.AutoMigrate(&models.User{})
}

func main() {
	r := gin.Default()

	r.GET("/", handlers.HomePageHandler)
	r.POST("/api/signup", handlers.SignupHandler)
	r.POST("/api/login", handlers.LoginHandler)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
