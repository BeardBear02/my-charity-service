package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//返回默认的路由引擎
	r := gin.Default()
	//r.GET("/hello", sayHello)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"bookname": "HarryPorter1",
		})
	})
	r.POST("/hello", func(c *gin.Context) {
		//http库里的OK就代表200
		c.JSON(http.StatusOK, gin.H{
			"bookname": "HarryPorter2",
		})
	})
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"bookname": "HarryPorter3",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"bookname": "HarryPorter4",
		})
	})

	//启动服务,没有这个run不了啊
	r.Run()
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello GoLang",
	})
}
