package main

import (
	"golang-eco/database"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	database.ConnectDatabase()
	router := gin.New()
	router.Use(gin.Logger())
	// routes.UserRoutes(router)
	// router.Use(middleware.Authentication)
	// router.GET("/addtocart", app.AddToCart())
	// router.GET("/removeitem", app.RemoveItem())
	// router.GET("/cartcheckout", app.BuyFromCart())
	// router.GET("/instantbuy", app.InstantBuy())

	// log.Fatal(router.Run(":" + port))
	router.Run(":8080")

}
