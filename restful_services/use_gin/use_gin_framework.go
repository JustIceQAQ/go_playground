package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/users", listUser)
	r.GET("/user/:id", retrieveUser)
	r.POST("/users", createUser)
	r.Run(":8000")
}

func listUser(c *gin.Context) {
	c.JSON(200, users)
}

func retrieveUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	found := false
	for _, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			user = u
			found = true
			break
		}
	}
	if found {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{
			"message": "沒有找到使用者",
		})
	}
}

func createUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name != "" {
		u := User{ID: len(users) + 1, Name: name}
		users = append(users, u)
		c.JSON(http.StatusCreated, u)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "缺少使用者名稱",
		})
	}
}

var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 3, Name: "王五"},
}

type User struct {
	ID   int
	Name string
}

// try POST
// windows: curl -Method 'POST' -Body "name=QAQ" -Uri 'http://127.0.0.1:8000/users'
// linux & Mac Os: curl -X POST -d 'name=QAQ' http://localhost:8080/users
