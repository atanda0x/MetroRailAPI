package controllers

import (
	"context"
	"time"

	"github.com/atanda0x/e-commerce/models"
	"github.com/gin-gonic/gin"
)

func HashPassword(password string) string {

}

func Verify(userPassword string, givenPassword string) (bool, string) {

}

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := ctx.BindJSON(&user); err != nil {

		}
	}
}

func Login() gin.HandlerFunc {

}

func ProductViewAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {

}
