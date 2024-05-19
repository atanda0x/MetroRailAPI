package routes

import "github.com/gin-gonic/gin"

func UserRoutes(inCommingRoute *gin.Engine) {
	inCommingRoute.POST("/users/signup")
	inCommingRoute.POST("/users/login")
	inCommingRoute.POST("/admin/addProduct")
	inCommingRoute.GET("/users/productView")
	inCommingRoute.GET("/users/search")
}
