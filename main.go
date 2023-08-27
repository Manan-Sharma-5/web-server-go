package main

import (
	"github.com/gin-gonic/gin"
	"web-server/Controller"
	"web-server/modal"
)

func main() {
	router := gin.Default()
	modal.ConnectDatabase()
	router.POST("/post", Controller.CreatePost)
	router.Run("localhost:8080")
}
