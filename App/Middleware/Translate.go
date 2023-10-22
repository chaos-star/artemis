package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	es_trans "github.com/go-playground/validator/v10/translations/es"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

func Translate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			langInst ut.Translator
			lang     = c.GetHeader("X-App-language")
			trans    = ut.New(zh.New(), en.New(), es.New())
		)

		langInst, _ = trans.FindTranslator(lang, "zh")
		switch lang {
		case "zh":
			_ = zh_trans.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), langInst)
		case "en":
			_ = en_trans.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), langInst)
		case "es":
			_ = es_trans.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), langInst)
		default:
			lang = "zh"
			_ = zh_trans.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), langInst)
		}
		c.Set("tran", langInst)
		fmt.Println(lang)
		c.Next()
	}
}
