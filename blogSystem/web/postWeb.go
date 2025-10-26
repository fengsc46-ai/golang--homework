package web

import (
	"blogSystem/bean"
	"blogSystem/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SavePost(c *gin.Context) {
	var post bean.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post create successfully"})
}

func QueryPostList(c *gin.Context) {
	var posts []bean.Post
	if err := db.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func PostDetail(c *gin.Context) {
	postId := c.Param("id")
	var post bean.Post
	if err := db.DB.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})

}

func DeletePost(c *gin.Context) {
	postId := c.Param("id")
	var post bean.Post
	if err := db.DB.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err := db.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Post delete failed"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post delete successfully"})
}
