package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	/*
		curl -X POST "http://127.0.0.1:8080/post" -d '{"name":"wang"}'
			{"name":"wang"}
		curl -X POST "http://127.0.0.1:8080/post" -d 'first_name=wang&last_name=kai'
			wang,kai,first_name=wang&last_name=kai
	*/
	r.POST("/post", func(c *gin.Context) {
		//读取出requestbody的字节
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		//如果错误不是空
		if err != nil {
			//返回并结束输出
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		//ioutil.ReadAll读取出数据后，原来数据则不存在，需要进行回存
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		//表单提交的格式post数据,如果不进行回存此处则无法拿到数据
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_last_name")
		//返回数据
		c.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(bodyBytes))
	})
	r.Run()
}
