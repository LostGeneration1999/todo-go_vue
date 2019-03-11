package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getTodoHandler(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)

	todo := getTodo(id)

	c.JSON(200, todo)
}

func getUserHandler(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)

	todo := getUser(id)

	c.JSON(200, todo)
}

func postUserHandler(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	user := registerUser(name, email)

	c.JSON(200, gin.H{
		"messege": "Success!",
		"result":  user,
	})

}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todo/:id", getTodoHandler)
	r.GET("/user/:id", getUserHandler)
	r.POST("/user", postUserHandler)

	r.Run()
}
