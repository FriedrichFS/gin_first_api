package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID         string `json:"id"`
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

func userById(c *gin.Context) {
	id := c.Param("id")
	user, err := getUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found!"})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}
func getUserById(id string) (*user, error) {
	for i, b := range user_list {
		if b.ID == id {
			return &user_list[i], nil
		}
	}
	return nil, errors.New("User not found!")
}

func createUser(c *gin.Context) {
	var newUser user
	var my_condition bool
	var error_list = []string{}

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	fmt.Println(newUser.First_name)
	for i, value := range user_list {
		if value.First_name == newUser.First_name {
			error_list = append(error_list, "First name already exists in DB!")
		}
		if value.Last_name == newUser.Last_name {
			error_list = append(error_list, "Last name already exists in DB!")
		}
		if value.Email == newUser.Email {
			error_list = append(error_list, "This Email already exists!")
		}
		fmt.Println(i)
	}

	if len(error_list) >= 0 {
		for counter, value := range user_list {
			fmt.Println("An exception occured! Error: ", value, counter)
		}
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "An error occured!  Check console for more information!"})
		return
	} else if my_condition {
		fmt.Println("The condition here is actually: ", my_condition)
		user_list = append(user_list, newUser)
		c.IndentedJSON(http.StatusCreated, newUser)
	}
}

func main() {
	router := gin.Default()
	router.GET("/get_users", getUsers)
	router.GET("/users/:id", userById)
	router.POST("/create_user", createUser)
	router.Run("localhost:8080")
}
