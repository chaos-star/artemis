package Response

import (
	"Artemis/Common/Code"
	"Artemis/Global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
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

func FailValidate(err error, c *gin.Context) {
	var (
		response Response
		lang     = GetLanguage(c)
		message  = ""
		errs     validator.ValidationErrors
	)
	response.Code = Code.ParamsException
	response.Message = Code.GetMessage(response.Code, lang)
	response.Data = gin.H{}

	if trans, ok := c.Get("tran"); ok && errors.As(err, &errs) {
		tranErrs := errs.Translate(trans.(ut.Translator))
		for _, val := range tranErrs {
			message = val
			break
		}
	} else {
		if err != nil {
			message = err.Error()
		}
	}
	response.Message = fmt.Sprintf("%s: %s", response.Message, message)
	c.JSON(http.StatusOK, response)
}

func GetLanguage(c *gin.Context) string {
	language := c.GetHeader("X-App-language")
	if language == "" {
		language = "zh"
	}
	return language
}

func Success(c *gin.Context) {
	var (
		response Response
		lang     = GetLanguage(c)
	)
	response.Code = Code.Success
	response.Message = Code.GetMessage(Code.Success, lang)
	response.Data = gin.H{}
	c.JSON(http.StatusOK, response)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	var (
		response Response
		lang     = GetLanguage(c)
	)
	response.Code = Code.Success
	response.Message = Code.GetMessage(Code.Success, lang)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

func Fail(code int, c *gin.Context) {
	var (
		response Response
		lang     = GetLanguage(c)
	)
	response.Code = code
	response.Message = Code.GetMessage(code, lang)
	response.Data = gin.H{}
	c.JSON(http.StatusOK, response)
}

func FailWithMessage(message string, c *gin.Context) {
	var (
		response Response
		lang     = GetLanguage(c)
	)
	response.Code = Code.BusinessException
	response.Message = Code.GetMessage(response.Code, lang)
	response.Message = fmt.Sprintf("%s: %s", response.Message, message)
	response.Data = gin.H{}
	c.JSON(http.StatusOK, response)
}

func FailWithData(code int, data interface{}, c *gin.Context) {
	var (
		response Response
		lang     = GetLanguage(c)
	)
	response.Code = code
	response.Message = Code.GetMessage(code, lang)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

func FailWithException(ex *Code.Exception, c *gin.Context) {
	var lang = GetLanguage(c)
	Fail(ex.Code(), c)
	log := Global.Log.GetLogger()
	log.SetReportCaller(false)

	fields := logrus.Fields{}
	fields["path"] = fmt.Sprintf("%s:%d", ex.FileName(), ex.Line())
	fields["method"] = ex.MethodName()
	message := Code.GetMessage(ex.Code(), lang)
	c.Set("_position_fields", fields)
	Global.Log.WithFields(fields).Error(fmt.Sprintf("Code:%d, Message:%s, Exception:%s", ex.Code(), message, ex.Error()))
	log.SetReportCaller(true)

}

func FailWithMessageAndData(message string, data interface{}, c *gin.Context) {
	var (
		response Response
		lang     = GetLanguage(c)
	)
	response.Code = Code.BusinessException
	response.Message = Code.GetMessage(response.Code, lang)
	response.Message = fmt.Sprintf("%s: %s", response.Message, message)
	response.Data = data
	c.JSON(http.StatusOK, response)
}
