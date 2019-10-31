package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get/:id/:name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id":   c.Param("id"),
			"name": c.Param("name"),
		})
	})

	r.Run()
}
