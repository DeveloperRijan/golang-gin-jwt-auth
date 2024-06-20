package main

import (
	"fmt"
	"golang_gin_jwt_auth/handlers"
	"golang_gin_jwt_auth/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.DBConnect()
	fmt.Println(initializers.DB.Config)
}

func main() {
	r := gin.Default()

	r.GET("/", handlers.HomePageHandler)
	r.POST("/api/login", handlers.LoginHandler)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
