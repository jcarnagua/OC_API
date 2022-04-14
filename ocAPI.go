package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func main() {
	router := gin.Default()
	hours, minutes, _ := time.Now().Clock()
	currTime := fmt.Sprintf("%02d%02d", hours, minutes)
	authorizedUser := user{Username: "c137@onecause.com", Password: "#th@nH@rm#y#r!$100%D0p#", Token: currTime}

	fmt.Println("Hello World! Current time is: " + currTime +
		"\nUsername is: " + authorizedUser.Username +
		"\nPassword is: " + authorizedUser.Password +
		"\nToken is: " + authorizedUser.Token)

	router.Run("localhost:8080")
}

func validateCredentials(c *gin.Context, authUser user) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	if newUser.Username == authUser.Username &&
		newUser.Password == authUser.Password &&
		newUser.Token == authUser.Token {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "authorized"})
		return
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	}

}
