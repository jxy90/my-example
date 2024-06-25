package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 定义一个结构体来存储 goroutine 的结果
type Result struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()

	// 启动一个 goroutine 持续运行

	// 定义一个处理函数来返回 goroutine 的结果
	r.GET("/result", func(c *gin.Context) {
		go func() {
			for {
				// 模拟一些持续运行的任务
				time.Sleep(5 * time.Second)
				log.Println("Goroutine is still running")
			}
		}()
		c.JSON(http.StatusOK, Result{Message: "return"})

	})

	// 启动 HTTP 服务器
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
