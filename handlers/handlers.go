package handlers

import (
	"fmt"
	"golang_gin_jwt_auth/helpers"
	"golang_gin_jwt_auth/models"

	"github.com/gin-gonic/gin"
)

// handlers
func HomePageHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Coming": "Soon",
	})
}

// Login
func LoginHandler(ctx *gin.Context) {
	var u models.User
	u.Email = ctx.PostForm("email")
	u.Password = ctx.PostForm("password")
	fmt.Println(fmt.Sprintf("Body: %s %s", u.Email, u.Password))

	//validations
	if u.Email == "" {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": "Email is required",
		})
		return
	}

	if !helpers.IsValidEmail(u.Email) {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": "Email address is invalid",
		})
		return
	}

	if u.Password == "" {
		ctx.JSON(400, gin.H{
			"success": false,
			"message": "Password is required",
		})
		return
	}

}
