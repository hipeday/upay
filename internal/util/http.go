package util

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func BuildQueryParams(payload interface{}) url.Values {
	params := url.Values{}
	v := reflect.ValueOf(payload)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := typeOfS.Field(i)
		queryTag := fieldType.Tag.Get("query")

		// 跳过没有 `query` 标签或包含 `skip` 的字段
		if queryTag == "" || containsSkip(queryTag) {
			continue
		}

		// 处理字段值
		if field.IsNil() {
			continue
		}

		queryParamName := getQueryParamName(queryTag)
		switch field.Kind() {
		case reflect.Ptr:
			elem := field.Elem()
			switch elem.Kind() {
			case reflect.String:
				params.Set(queryParamName, elem.String())
			case reflect.Bool:
				params.Set(queryParamName, strconv.FormatBool(elem.Bool()))
			case reflect.Int32:
				params.Set(queryParamName, strconv.Itoa(int(elem.Int())))
			default:
				panic("unhandled default case")
			}
		}
	}

	return params
}

// 辅助函数：检查标签中是否包含 `skip`
func containsSkip(tag string) bool {
	parts := splitTag(tag)
	for _, part := range parts {
		if part == "skip" {
			return true
		}
	}
	return false
}

// 辅助函数：获取 `query` 标签的第一个部分作为查询参数名称
func getQueryParamName(tag string) string {
	parts := splitTag(tag)
	return parts[0]
}

// 辅助函数：拆分标签
func splitTag(tag string) []string {
	return strings.Split(tag, ",")
}
