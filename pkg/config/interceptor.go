package config

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// http请求 自定义拦截器
func HttpInterceptor() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Println("用时:", latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println("请求状态:", status)

	}

}
