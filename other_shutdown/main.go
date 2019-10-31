package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
测试流程：
	命令启动服务：go run 以便于测试ctrl+c关闭服务器
	访问接口后，立即返回服务器端，ctrl+c关闭服务器
	看最后服务端的输出结果和客户端的响应结果
*/
func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(200, "hello   test")
	})

	//创建服务
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	//将服务放在协程里
	go func() {
		//创建监听服务，并返回错误，如果错误不是空并且不是关闭服务器，即正常监听时，打印错误
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//真正的请求拦截，放在channl里面，类型是os.Signal
	quit := make(chan os.Signal)
	//下面是signal退出信号捕获
	//传入channel
	//主要捕获两类：kill -9 无法捕获强制退出
	//捕获ctrl+c和kill -15 的终止
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//channel阻塞
	<-quit
	//日志打印
	log.Println("shutdown server ...")
	//创建超时上下文,后面定义的时间为关闭服务器最大的等待时间，，如果响应时间超过这个时间则会强制关闭
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown ： ", err)
	}
	log.Println("server exiting ...")
}
