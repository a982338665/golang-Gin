package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//路由路径 根路径
	r.Static("/assets", "./router_static/assets")
	r.StaticFS("/static", http.Dir("./router_static/static"))
	r.StaticFile("/favicon.ico", "./router_static/favicon.ico")
	r.Run(":8000")
}
