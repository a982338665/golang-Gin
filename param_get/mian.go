package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		//取？参数
		firstName := c.Query("first_name")
		//取？参数,若没有则返回默认值adminUser
		lastName := c.DefaultQuery("last_name", "adminUser")
		c.String(http.StatusOK, "%s,%s", firstName, lastName)
	})
	r.Run(":8080")

}
