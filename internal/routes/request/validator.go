package request

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	errors2 "github.com/hipeday/upay/internal/errors"
	"reflect"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Validatable interface {
	Validate() error
}

func Validate(s Validatable) error {
	err := validate.Struct(s)
	if err != nil {
		var (
			errs    validator.ValidationErrors
			message string
		)
		if errors.As(err, &errs) {
			for _, e := range errs {
				// 通过反射获取字段
				field, ok := reflect.TypeOf(s).FieldByName(e.Field())
				if !ok {
					continue
				}
				// 获取message标签的值
				message = field.Tag.Get("message")
				if message == "" {
					// 如果没有提供message，使用默认的错误描述
					message = fmt.Sprintf("%s is invalid", e.Field())
				}
			}
			// 合并所有错误消息为一个字符串返回
			return errors2.NewInternalServerError(message)
		}
	}
	return nil
}
