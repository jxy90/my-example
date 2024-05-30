package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

const XRequestID = "X-Request-Id"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 X-Request-Id
		headerKey := XRequestID
		val := c.GetHeader(headerKey)
		if val == "" {
			// 如果没有 X-Request-Id，则生成一个新的 UUID
			val = uuid.New().String()
			// 设置新的 X-Request-Id 到请求头中
			c.Request.Header.Set(headerKey, val)
		}
		// 将 X-Request-Id 设置到响应头中
		c.Header(headerKey, val)
		c.Next()
	}
}
func PRequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 获取请求头中的 X-Request-Id
		headerKey := XRequestID
		val := c.GetHeader(headerKey)
		log.Println("traceId:" + val)
	}
}

func main() {
	r := gin.Default()

	// 使用中间件
	r.Use(RequestIDMiddleware(), PRequestIDMiddleware())

	r.GET("/", func(c *gin.Context) {
		// 读取 X-Request-Id
		requestID := c.GetHeader(XRequestID)
		c.JSON(200, gin.H{
			"message":    "Hello World",
			"request_id": requestID,
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
