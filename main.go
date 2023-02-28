package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

type user struct {
	ID         string `json:"ID"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

var user_list = []user{
	{ID: "1", First_name: "Friedrich", Last_name: "Rösel", Email: "friedrich2208@gmail.com", Password: "Test_PW"},
	{ID: "2", First_name: "Hans", Last_name: "Wurst", Email: "hans@wurst.com", Password: "Test_PW22"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user_list)
}

func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	user_list = append(user_list, newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.GET("/get_users", getUsers)
	router.POST("/create_user", createUser)
	router.Run("localhost:8080")
}
