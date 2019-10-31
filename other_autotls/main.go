package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

//访问时：https://...
//需要外网部署测试
func main() {

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	//自动证书下载
	//1.生成本地秘钥
	//2.发送秘钥给证书颁发机构，获取一个私钥
	//3.拿到私钥进行验证，之后请求通过私钥加密
	autotls.Run(r, "www.itpp.tk")

}
