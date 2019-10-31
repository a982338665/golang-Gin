package main

import "github.com/gin-gonic/gin"

//http://localhost:8080/index
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./other_template/template/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "index.html",
		})
	})
	r.Run()
}
