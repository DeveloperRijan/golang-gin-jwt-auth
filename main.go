package main

import (
	"fmt"
	"golang_gin_jwt_auth/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

// handlers
func HomePageHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Coming": "Soon",
	})
}

func init() {
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()

	r.GET("/", HomePageHandler)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
