package main

import (
	"context"
	"errors"
	"fmt"
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
	r.GET("/", defaultH)
	r.GET("/result", handler)
	r.GET("/result2", handler2)

	// 启动 HTTP 服务器
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

// 默认测试handler
func defaultH(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			fmt.Println("err defer")
		} else {
			fmt.Println("nil defer")
		}
	}()
	err = errors.New("err New")
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, Result{Message: "return"})
}

// context结束goroutine不结束
func handler(c *gin.Context) {
	go func(c *gin.Context) {
		for {
			// 模拟一些持续运行的任务
			time.Sleep(1 * time.Second)
			log.Println("Goroutine is still running")
		}
	}(c)
	c.JSON(http.StatusOK, Result{Message: "return"})
}

// context结束goroutine结束
func handler2(c *gin.Context) {
	// 创建一个带取消功能的上下文
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel() // 确保在函数退出时取消上下文
	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				log.Println("Goroutine is stopped")
				return
			default:
				// 模拟一些持续运行的任务
				time.Sleep(1 * time.Second)
				log.Println("Goroutine is still running")
			}
		}
	}(ctx)
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, Result{Message: "return2"})
}
