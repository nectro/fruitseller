package handlers

import (
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProdToCart(c *gin.Context) {
	token := c.GetHeader("key")

	var payload ProductToCartRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify the auth token
	user := services.VerifyAuthToken(token)

	if user != "" {
		cartid := services.GetUserDetails(user, "CARTID")
		check := services.UpdateCartAddProd(cartid, payload.ProdId)
		c.JSON(http.StatusOK, gin.H{
			"Success": check,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
		})
	}
}

func RemoveProdFromCart(c *gin.Context) {
	token := c.GetHeader("key")

	var payload ProductToCartRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify the auth token
	user := services.VerifyAuthToken(token)

	if user != "" {
		cartid := services.GetUserDetails(user, "CARTID")
		check := services.UpdateCartRemoveProd(cartid, payload.ProdId)
		c.JSON(http.StatusOK, gin.H{
			"Success": check,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
		})
	}
}

func ReadProdFromCart(c *gin.Context) {
	token := c.GetHeader("key")

	// verify the auth token
	user := services.VerifyAuthToken(token)

	if user != "" {
		cartid := services.GetUserDetails(user, "CARTID")
		ProductList := services.ReadProductsFromCart(cartid)
		c.JSON(http.StatusOK, gin.H{
			"Success": true,
			"data":    ProductList,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
		})
	}
}

func CartDetails(c *gin.Context) {
	token := c.GetHeader("key")

	// verify the auth token
	user := services.VerifyAuthToken(token)

	if user != "" {
		cartid := services.GetUserDetails(user, "CARTID")
		ProductList := services.ReadProductsFromCart(cartid)
		c.JSON(http.StatusOK, gin.H{
			"Success":     true,
			"CartId":      cartid,
			"ProductList": ProductList,
			"Total Items": len(ProductList),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
		})
	}
}
