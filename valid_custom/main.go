package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

type Booking struct {
	//登录时间:加自定义校验器bookableDate
	CheckIn time.Time `form:"check_in" binding:"required,bookableDate11" time_format:"2006-01-02"`
	//登出时间，由于登出时间大于登入时间，所以此处要加验证 gtfield=checkIn
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=checkIn" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	//注册校验器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//第一个参数为结构体中
		//第二个参数为方法
		v.RegisterValidation("bookableDate11", bookableDate)
	}
	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(b); err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "ok",
			"booking": b,
		})
	})
}

//引入自定义
// Structure
func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	//如果能拿到开始时间
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		//开始时间大于今天，就可以接受预定，校验通过
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return true
}
