package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getTodoHandler(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)

	todo, err := getTodo(id)
	if err != nil {
		c.JSON(404, gin.H{
			"messege": "Not Found Todo!",
		})
		return
	}
	c.JSON(200, todo)
}

func getUserHandler(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)

	user, err := getUser(id)
	if err != nil {
		c.JSON(404, gin.H{
			"messege": "Not Found User!",
		})
		return
	}
	c.JSON(200, user)
}

func postUserHandler(c *gin.Context) {
	email := c.PostForm("email")
	passoword := c.PostForm("password")

	user, err := registerUser(email, passoword)
	if err != nil {
		c.JSON(404, gin.H{
			"messege": err,
		})
		return
	}

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
