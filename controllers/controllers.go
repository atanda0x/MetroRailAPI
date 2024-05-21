package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/atanda0x/e-commerce/database"
	"github.com/atanda0x/e-commerce/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cancel context.CancelFunc
		var ctx context.Context
		ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		var count int
		err := database.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE email = $1", user.Email).Scan(&count)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		passwordString := string(passwordHash)
		user.Password = &passwordString
		user.ID = uuid.New()
		user.CreatedAt = time.Now().UTC()
		user.UpdatedAt = time.Now().UTC()

		query := `
			INSERT INTO users (id, first_name, last_name, password, email, phone, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`

		_, err = database.DB.ExecContext(ctx, query, user.ID, user.First_Name, user.Last_Name, user.Password, user.Email, user.Phone, user.CreatedAt, user.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		c.JSON(http.StatusCreated, "Successfully Signed Up!")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement login logic here
	}
}

func ProductViewAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement product view admin logic here
	}
}

func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement search product logic here
	}
}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement search product by query logic here
	}
}
