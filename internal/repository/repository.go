package repository

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	insertSqlTemplate = "INSERT INTO %s (%s) VALUES (%s)"
)

func getColumns(obj interface{}) []string {
	// 获取对象的反射类型
	val := reflect.ValueOf(obj)

	// 如果传入的是指针，获取指针指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var tags []string

	// 遍历结构体的字段
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		// 如果字段是嵌入的结构体，递归获取嵌入结构体的字段标签
		if field.Anonymous {
			tags = append(tags, getColumns(val.Field(i).Interface())...)
		} else {
			// 获取 `db` 标签的值
			tag := field.Tag.Get("db")
			if tag != "" {
				tags = append(tags, tag)
			}
		}
	}

	return tags
}

func getInsertSql(tableName, columns string, columnCount int) string {
	var placeholder []string
	for i := 0; i < columnCount; i++ {
		placeholder = append(placeholder, "?")
	}
	return fmt.Sprintf(insertSqlTemplate, tableName, columns, strings.Join(placeholder, ","))
}

func getUpdateSql(tableName string, updateParameters []string, whereParameters []string) string {
	var (
		updates   []string
		wheres    []string
		sourceSql = "UPDATE %s SET %s"
	)
	if len(updateParameters) > 0 {
		for _, param := range updateParameters {
			updates = append(updates, param+" = ?")
		}
	}

	if len(whereParameters) > 0 {
		sourceSql += " WHERE %s"
		for _, param := range whereParameters {
			wheres = append(wheres, param+" = ?")
		}
	}
	return fmt.Sprintf(sourceSql, tableName, strings.Join(updates, ", "), strings.Join(wheres, " AND "))
}

func getQuerySql(tableName, columns string, parameters []string) string {
	var wheres []string
	if len(parameters) > 0 {
		for _, param := range parameters {
			wheres = append(wheres, param+" = ?")
		}
	}
	return fmt.Sprintf("SELECT %s FROM %s WHERE %s", columns, tableName, strings.Join(wheres, " AND "))
}
