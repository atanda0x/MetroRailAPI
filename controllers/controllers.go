package controllers

import (
	"context"
	"net/http"
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := userCollection.CountDoument(ctx, json.)
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
