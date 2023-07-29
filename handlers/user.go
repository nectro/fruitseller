package handlers

import (
	"backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func AddUser(c *gin.Context) {

	var payload AddRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if payload.Firstname != "" && payload.Lastname != "" && payload.Password != "" && payload.Email != "" && payload.Role != "" {
		generatedId := services.AddCart()

		if generatedId != uuid.Nil {
			services.AddUser(payload.Firstname, payload.Lastname, payload.Password, payload.Email, payload.Role, generatedId)

			c.JSON(http.StatusOK, gin.H{
				"Success": true,
				"data":    payload,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Success": false,
				"message": "Something went wrong!",
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"message": "add the data are not present!",
		})
	}

}

func LoginUser(c *gin.Context) {
	var payload LoginRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Success": false, "error": err.Error()})
		return
	}

	data := services.AuthenticateUser(payload.Email, payload.Password)

	if data != "failed" {
		checkOk := services.AddAuth(data, payload.Email)
		if checkOk == true {
			c.JSON(http.StatusOK, gin.H{"Success": true, "token": data})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Success": false, "message": "couldn't add auth token!"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Success": false, "message": "couldn't generate auth token!"})
	}
}

// "firstname":"root1",
// "lastname":"admin",
// "email":"demo1@gmail.com",
// "role":"ADMIN",
// "password":"dasd"
