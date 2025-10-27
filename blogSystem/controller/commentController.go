package controller

import (
	"blogSystem/bean"
	"blogSystem/initial"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
}

type CommentRequest struct {
	Id          uint   `json:"id"`
	PostID      uint   `json:"postID"`
	UserID      uint   `json:"userID"`
	ContentText string `json:"contentText"`
}

func (controller *CommentController) SaveComment(c *gin.Context) {
	var request CommentRequest

	// 待补充对postId的合法性校验

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := bean.Comment{
		PostID:      request.PostID,
		UserID:      request.UserID,
		ContentText: request.ContentText,
	}
	if err := initial.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Comment created successfully"})
}

func (controller *CommentController) QueryCommentsByPostId(c *gin.Context) {
	postId := c.Param("postId")

	// 待补充对postId的合法性校验
	// 待补充对postId的存在性校验

	var request CommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var comments []bean.Comment
	if err := initial.DB.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query comments"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}
