package Response

import (
	"Artemis/Common/Code"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageData struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	List     interface{} `json:"list"`
}

type ListData struct {
	List interface{} `json:"list"`
}

func Success(c *gin.Context) {
	var response Response
	response.Code = Code.Success
	response.Message = Code.GetMessage(Code.Success)
	response.Data = gin.H{}
	c.JSON(http.StatusOK, response)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	var response Response
	response.Code = Code.Success
	response.Message = Code.GetMessage(Code.Success)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

func Fail(code int, c *gin.Context) {
	var response Response
	response.Code = code
	response.Message = Code.GetMessage(code)
	response.Data = gin.H{}
	c.JSON(http.StatusOK, response)
}

func FailWithData(code int, data interface{}, c *gin.Context) {
	var response Response
	response.Code = code
	response.Message = Code.GetMessage(code)
	response.Data = data
	c.JSON(http.StatusOK, response)
}