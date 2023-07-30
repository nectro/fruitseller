package handlers

import (
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProducts(c *gin.Context) {
	token := c.GetHeader("key")

	var payload ProductAddRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify the auth token
	user := services.VerifyAuthToken(token)

	if user != "" {

		// get the user role details
		role := services.GetUserDetails(user, "ROLE")

		if role == "ADMIN" {
			// after adding the product get the product id
			prodId := services.AddProduct(payload.Name, payload.Price, payload.ImageId, payload.Description)
			c.JSON(http.StatusOK, gin.H{
				"Success": true,
				"email":   user,
				"role":    role,
				"prodId":  prodId,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Success": false,
				"message": "Access Denied!",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"email":   user,
		})
	}
}

func DeleteProducts(c *gin.Context) {
	token := c.GetHeader("key")

	var payload ProductDeleteRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify the auth token
	user := services.VerifyAuthToken(token)

	if user != "" {

		// get the user role details
		role := services.GetUserDetails(user, "ROLE")

		if role == "ADMIN" {
			// after adding the product get the product id
			check := services.DeleteProduct(payload.Id)
			c.JSON(http.StatusOK, gin.H{
				"Success": check,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Success": false,
				"message": "Access Denied!",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"email":   user,
		})
	}
}

func UpdateProduct(c *gin.Context) {
	token := c.GetHeader("key")

	var payload ProductUpdateRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify the auth token
	user := services.VerifyAuthToken(token)

	if user != "" {

		// get the user role details
		role := services.GetUserDetails(user, "ROLE")

		if role == "ADMIN" {
			// after adding the product get the product id
			check := services.UpdateProduct(payload.Id, payload.Name, payload.Price, payload.ImageId, payload.Description)
			c.JSON(http.StatusOK, gin.H{
				"Success": check,
				"email":   user,
				"role":    role,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Success": false,
				"message": "Access Denied!",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"message": "first login",
		})
	}
}

func ReadProducts(c *gin.Context) {
	var responseArray []services.Product

	responseArray = services.ReadProducts()

	if responseArray != nil {
		c.JSON(http.StatusOK, gin.H{
			"Success": true,
			"data":    responseArray,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
		})
	}
}

func ReadProduct(c *gin.Context) {
	prodId := c.Param("id")

	var response services.Product
	var UNDEFINED services.Product

	response = services.ReadProduct(prodId)

	if response != UNDEFINED {
		c.JSON(http.StatusOK, gin.H{
			"Success": true,
			"data":    response,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
		})
	}
}
