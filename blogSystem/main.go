package main

import (
	"blogSystem/api"
	"blogSystem/initial"

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
	initial.InitDbConnection()
	// Initialize api routes
	initWebRoute()

}

func initWebRoute() {
	// 创建路由
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	{
		// 无需认证的接口分组
		uAuthGroup := r.Group("/uAuthGroup")
		// 首页路由
		uAuthGroup.GET("/", func(c *gin.Context) {
			c.String(200, "hello world!")
		})
		// 注册用户路由
		uAuthGroup.POST("/register", api.Register)
		// 登录用户路由
		uAuthGroup.POST("/login", api.Login)
	}

	{
		//需要认证的接口分组
		authGroup := r.Group("/auth")
		authGroup.Use(api.AuthMiddleware())
		// 文章相关路由
		authGroup.GET("/postList", api.QueryPostList)
		authGroup.GET("/postDetail/:id", api.PostDetail)
		authGroup.POST("savePost", api.SavePost)
		authGroup.DELETE("/deletePost/:id", api.DeletePost)

		// 评论相关路由
		authGroup.POST("/saveComment", api.SaveComment)
		authGroup.POST("/queryCommentsByPostId/:postId", api.QueryCommentsByPostId)
	}
	// 启动服务器
	err := r.Run(":9000")
	if err != nil {
		panic(err)
	}
}
