package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	// 创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./favicon.ico"))
	// 连接数据库

	// 加载静态网页
	ginServer.LoadHTMLGlob("./templates/*")

	// 加载资源文件
	ginServer.Static("/static", "./static")

	// 响应一个静态网页
	ginServer.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是go后台传回来的数据",
		})
	})

	// 访问地址 处理请求
	ginServer.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// Get请求
	ginServer.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// Post请求
	ginServer.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "post,user",
		})
	})

	// 第一种请求传参 Query(key)
	ginServer.GET("/user/info", func(c *gin.Context) {
		userId := c.Query("userId")
		userName := c.Query("userName")
		c.JSON(http.StatusOK, gin.H{
			"userId":   userId,
			"userName": userName,
		})
	})

	// 第二种请求传参
	ginServer.GET("/user/info1/:userId/:userName", func(c *gin.Context) {
		userId := c.Param("userId")
		userName := c.Param("userName")
		c.JSON(http.StatusOK, gin.H{
			"userId":   userId,
			"userName": userName,
		})
	})

	// 前端给后端传递Post JSON
	ginServer.POST("/json", func(c *gin.Context) {
		b, _ := c.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})

	// 表单
	ginServer.POST("/user/add", func(c *gin.Context) {
		userName := c.PostForm("userName")
		userId := c.PostForm("userId")
		c.JSON(http.StatusOK, gin.H{
			"userName": userName,
			"userId":   userId,
		})
	})

	// 路由
	ginServer.GET("/test", func(c *gin.Context) {
		// 重定向
		c.Redirect(301, "http://www.baidu.com")
	})
	ginServer.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
