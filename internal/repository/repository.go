package repository

import (
	"fmt"
	"strings"
)

const (
	insertSqlTemplate = "INSERT INTO %s (%s) VALUES (%s)"
)

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
