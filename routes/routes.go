package routes

import (
	"golang-eco/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.POST("/admin/addproduct", controller.AddProduct())
	incomingRoutes.GET("/users/productview", controller.ProductViewerAdmin())
	incomingRoutes.GET("/users/search", controller.SearchProduct())
}
