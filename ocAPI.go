package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Token      string `json:"token"`
	Authorized bool   `json:"authorized"`
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "World"})
	})

	router.POST("/login", validateCredentials)

	router.Run(":8080")
}

func validateCredentials(c *gin.Context) {
	var newUser User
	hours, minutes, _ := time.Now().Clock()
	currTime := fmt.Sprintf("%02d%02d", hours, minutes)
	authorizedUser := User{Username: "c137@onecause.com", Password: "#th@nH@rm#y#r!$100%D0p#", Token: currTime}

	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.Authorized = false

	if newUser.Username == authorizedUser.Username &&
		newUser.Password == authorizedUser.Password &&
		newUser.Token == authorizedUser.Token {
		newUser.Authorized = true
		c.JSON(http.StatusOK, gin.H{"authorized": newUser.Authorized})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authorized": newUser.Authorized})

}
