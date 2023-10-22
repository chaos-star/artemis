package Middleware

import (
	"Artemis/Global"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chaos-star/marvel"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

func TraceLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		Url := c.Request.URL.String()
		Method := c.Request.Method
		reqHeader := c.Request.Header.Clone()
		Header, _ := json.Marshal(reqHeader)
		reqBody, _ := io.ReadAll(c.Request.Body)
		reqTpl := fmt.Sprintf("Url: %s | Method:%s | Header:%s | Params:%s", Url, Method, string(Header), string(reqBody))
		//request log
		Global.Log.Info(fmt.Sprintf("[Api-Request] %s", reqTpl))
		c.Request.Body = io.NopCloser(bytes.NewReader(reqBody))
		proxy := &responseProxy{ResponseWriter: c.Writer}
		c.Writer = proxy
		c.Next()
		//response log
		status := proxy.status
		body := proxy.body
		if position, has := c.Get("_position_fields"); has {
			if fields, ok := position.(logrus.Fields); ok {
				log := Global.Log.GetLogger()
				log.SetReportCaller(false)
				Global.Log.WithFields(fields).Info(fmt.Sprintf("[Api-Response] %s, HttpStatus:%d, Response:%s", reqTpl, status, string(body)))
				log.SetReportCaller(true)
				return
			}
		}
		marvel.Logger.Info(fmt.Sprintf("[Api-Response] %s, HttpStatus:%d, Response:%s", reqTpl, status, string(body)))
		// 在控制台中记录响应结果
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
