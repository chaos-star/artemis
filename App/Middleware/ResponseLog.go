package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ResponseRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 让请求继续流经中间件链
		proxy := &responseProxy{ResponseWriter: c.Writer}
		c.Writer = proxy
		c.Next()

		// 在请求处理完毕后，获取响应状态码和内容
		status := proxy.status
		body := proxy.body

		// 在控制台中记录响应结果
		fmt.Printf("Response Status: %d\n", status)
		fmt.Printf("Response Body: %s\n", body)
	}
}

type responseProxy struct {
	gin.ResponseWriter
	status int
	body   []byte
}

func (r *responseProxy) Write(data []byte) (int, error) {
	r.body = data
	return r.ResponseWriter.Write(data)
}

func (r *responseProxy) WriteString(s string) (int, error) {
	r.body = []byte(s)
	return r.ResponseWriter.WriteString(s)
}
