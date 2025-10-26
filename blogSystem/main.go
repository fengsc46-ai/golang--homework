package main

import (
	"blogSystem/db"
	"blogSystem/web"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: feng
 * @Date: 2021/11/26 16:30
 * @Description:
=====主要功能===========
1。登录、注册功能
2.文章的增删查改功能
3.评论的创建、查询功能

=====系统架构模块========
1.错误码统一处理，按类型返回
2.使用日志库记录运行信息和错误信息
*/

func main() {
	// Initialize database connection
	db.InitDbConnection()
	// Initialize web routes
	initWebRoute()

}

func initWebRoute() {
	// 创建路由
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 注册用户路由
	r.POST("/register", web.Register)
	// 登录用户路由
	r.POST("/login", web.Login)

	// 文章相关路由
	r.GET("/postList", web.QueryPostList)
	r.GET("/postdetail/:id", web.PostDetail)
	r.POST("savePost", web.SavePost)
	r.DELETE("/deletePost/:id", web.DeletePost)

	// 评论相关路由
	r.POST("/saveComment", web.SaveComment)
	r.POST("/queryCommentsByPostId/:postId", web.QueryCommentsByPostId)
	// 启动服务器
	err := r.Run("9000")
	if err != nil {
		panic(err)
	}
}
