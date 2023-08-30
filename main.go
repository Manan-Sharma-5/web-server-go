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
	router.GET("/post", Controller.ReadPost)
	router.GET("/post/:id", Controller.FindPost)
	router.PATCH("/post/:id", Controller.UpdatePost)
	router.DELETE("/post/:id", Controller.DeletePost)
	router.Run("localhost:8080")
}
