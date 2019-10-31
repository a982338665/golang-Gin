package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

//binging换成validate
type Person struct {
	Age     int    `form:"age" validate:"required,gt=10"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

/*
	测试访问：
http://localhost:8080/testing	 		[Age为必填字段 Name为必填字段 Address为必填字段]
http://localhost:8080/testing?age=18	[Name为必填字段 Address为必填字段]
http://localhost:8080/testing?age=18&locale=en	[Name is a required field Address is a required field]
*/
//验证信息多语言化
func main() {
	//创建验证器
	Validate := validator.New()
	zh := zh2.New()
	en := en2.New()
	//创建翻译器，里面设置和支持的语言为 zh 和 en
	Uni := ut.New(zh, en)

	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		//根据请求获取相应的翻译器
		trans, _ := Uni.GetTranslator(locale)

		//注册匹配到的验证器
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		}

		var person Person
		//绑定
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}
		//验证
		if err := Validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(500, "%v", sliceErrs)
			c.Abort()
			return
		}

	})
	r.Run()
}
