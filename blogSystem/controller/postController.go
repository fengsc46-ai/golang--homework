package controller

import (
	"blogSystem/bean"
	"blogSystem/initial"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

type PostRequest struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userID"`
}

func (controller *PostController) SavePost(c *gin.Context) {
	var request PostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := bean.Post{
		Title:   request.Title,
		Content: request.Content,
		UserID:  request.UserID,
	}
	if err := initial.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post create successfully"})
}

func (controller *PostController) QueryPostList(c *gin.Context) {
	var posts []bean.Post
	if err := initial.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (controller *PostController) PostDetail(c *gin.Context) {
	postId := c.Param("id")
	var post bean.Post
	if err := initial.DB.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"post": post})

}

func (controller *PostController) DeletePost(c *gin.Context) {
	postId := c.Param("id")

	//代补充对postId的合法性校验
	//待补充对post和user的存在校验

	var post bean.Post
	if err := initial.DB.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err := initial.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Post delete failed"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post delete successfully"})
}
