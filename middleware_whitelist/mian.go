package main

import "github.com/gin-gonic/gin"

func IPAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//访问者ip
		ip := c.ClientIP()
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		for _, host := range ipList {
			if ip == host {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "%s , not in ipList ", ip)
			c.Abort()
		}

	}
}

//自定义中间件，参考logger和recovery
//测试：http://127.0.0.1:8080/test
//测试：http://localhost:8080/test
//测试：http://192.168.31.60:8080/test 192.168.31.60 , not in ipList
func main() {

	r := gin.Default()
	r.Use(IPAuthMiddleWare())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello,test")
	})
	r.Run()
}
