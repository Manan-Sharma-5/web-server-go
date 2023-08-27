package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-server/modal"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	post := modal.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	modal.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}
