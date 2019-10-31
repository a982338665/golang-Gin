package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	//定义日志文件输出路径，在项目路径下
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Logger()) //申明中间件，原来使用的	r := gin.Default()，点进去可以看到默认使用了两个中间件，此处单独挑出来使用
	//r.Use(gin.Logger(),gin.Recovery())//gin.Recovery()中间件用来处理panic异常防止服务器挂掉
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		//panic("test panic")
		c.String(200, "%s", name)
	})

	r.Run()

}
