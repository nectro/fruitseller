package main

import (
	"backend/handlers"
	"backend/services"
	"github.com/gin-gonic/gin"
)

func main() {
	services.DbConnect() // initiate the connection

	defer services.DbClose() // close the connection once the server is terminated

	r := gin.Default()

	// Authentication routes
	r.POST("/user/add", handlers.AddUser)
	r.POST("/user/login", handlers.LoginUser)

	// Seller only actions
	r.POST("/products/add", handlers.AddProducts)
	r.POST("/products/delete", handlers.DeleteProducts)
	r.POST("/products/update", handlers.UpdateProduct)

	// buyer action
	r.POST("/cart/addProduct", handlers.AddProdToCart)
	r.POST("/cart/removeProduct", handlers.RemoveProdFromCart)
	r.GET("/cart", handlers.ReadProdFromCart)
	r.GET("/cart/details", handlers.CartDetails)

	// Both Buyer and seller privileges
	r.GET("/products", handlers.ReadProducts)
	r.GET("/products/:id", handlers.ReadProduct)

	r.Run() // port 0.0.0.0:8080
}
