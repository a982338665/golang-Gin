package main

import (
	"github.com/gin-gonic/gin"
)

//form表示可以由form转换为结构体,可以加参数验证,也包含一些字符串的验证，是否仅包含数字，或字母等
//binding:"required,gt=10" 表示必传并且要大于10
//binding:"required|gt=10" 表示必传或者要大于10
type Person struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
	age     int    `form:"age" binding:"required,gt=10"`
}

//请求参数验证：文档：https://godoc.org.gopkg.in/go-playground/validator.v8
func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		if error := c.ShouldBind(&person); error != nil {
			c.String(500, "%v", error)
			c.Abort()
			return
		}
		c.String(200, "%v", person)
	})
}
