package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//所有以/user为前缀的都会经过此路由
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run()
}
