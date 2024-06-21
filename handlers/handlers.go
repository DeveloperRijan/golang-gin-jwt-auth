package handlers

import (
	"fmt"
	"golang_gin_jwt_auth/helpers"
	"golang_gin_jwt_auth/initializers"
	"golang_gin_jwt_auth/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	//Find the user
	user := initializers.DB.First(&u, "email = ?", fmt.Sprintf("%s", u.Email))
	fmt.Println(user)
	// if user.Error != nil {
	// 	ctx.JSON(404, gin.H{
	// 		"success": false,
	// 		"message": fmt.Sprintf("Error! %s", user.Error),
	// 	})
	// 	return
	// }

	//generate jwt
	var (
		jwt_key       string
		jwt_token     *jwt.Token
		new_jwt_token string
		err           error
	)
	jwt_key = os.Getenv("JWT_KEY")
	jwt_token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	new_jwt_token, err = jwt_token.SignedString([]byte(jwt_key))

	if err != nil {
		println(err.Error())
		ctx.JSON(500, gin.H{
			"success": false,
			"message": "Generating token failed, please try again!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"success": true,
		"token":   new_jwt_token,
	})
	return
}
