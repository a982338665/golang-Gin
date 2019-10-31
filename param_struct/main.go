package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

//form表示可以由form转换为结构体
type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

/*
curl -X GET 'http://127.0.0.1:8080/testing?name=wang&address=wuhan&birthday=2000-01-06'
curl -X POST 'http://127.0.0.1:8080/testing?name=wang&address=wuhan&birthday=2000-01-06'
curl -X POST 'http://127.0.0.1:8080/testing' -d 'name=wang&address=wuhan&birthday=2000-01-06'
curl -H 'Content-Type:application/json' -X POST 'http://127.0.0.1:8080/testing' -d '{"name":"wang"}'
*/
func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
}

func testing(c *gin.Context) {
	var person Person
	//此处是根据请求的content-type来区分匹配，做不同的绑定操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(500, "person bind error:%v", err)
	}
}
