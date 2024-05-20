package main

import (
	"log"
	"os"

	"github.com/atanda0x/e-commerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use()

	router.GET("/addToCart")
	router.GET("/removeItem")
	router.GET("/cartCheckout")
	router.GET("/instantBuy")

	log.Fatal(router.Run(":", port))
}
