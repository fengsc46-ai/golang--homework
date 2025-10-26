package api

import (
	"blogSystem/bean"
	"blogSystem/initial"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveComment(c *gin.Context) {
	var comment bean.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initial.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Comment created successfully"})
}

func QueryCommentsByPostId(c *gin.Context) {
	var comments []bean.Comment
	postId := c.Param("postId")
	if err := initial.DB.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query comments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}
