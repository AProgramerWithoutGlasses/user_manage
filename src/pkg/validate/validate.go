package validate

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

func Init() (err error) {
	// 注册自定义验证器
	err = validate.RegisterValidation("no-whitespace", noWhitespace)
	if err != nil {
		fmt.Println("validate.RegisterValidation() err", err)
		return
	}

	return
}

// Validation 用于参数校验
func Validate(param any) (err error) {
	err = validate.Struct(param)
	if err != nil {
		fmt.Println("validate.Struct() err", err)
		err = tidyErrMsg(err) // 整理错误信息
		return
	}
	return
}

// noWhitespace 自定义验证器：检查是否包含空格
func noWhitespace(fl validator.FieldLevel) bool {
	return !strings.Contains(fl.Field().String(), " ")
}

// tidyErrMsg 将错误字段和错误提示整理拼接并返回
func tidyErrMsg(err error) (err2 error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			field := fieldMsg[vErr.StructField()]
			errMes := errMsg[vErr.Tag()]
			fmt.Println(vErr.StructField(), vErr.Tag())
			return errors.New(field + errMes)
		}
	}

	return
}
