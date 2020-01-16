package generate

import (
	"fmt"
	"fzw/fxxkModel/conf"
	"fzw/fxxkModel/tools"
	"io"
	"os"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/generator"
)

/*
	生成model结构体
*/
func Generate(tableNames ...string) {
	// 生成表信息
	tableNamesStr := ""
	for _, name := range tableNames {
		if tableNamesStr != "" {
			tableNamesStr += ","
		}
		tableNamesStr += "'" + name + "'"
	}
	tables := getTables(tableNamesStr)
	// 生成模型
	for _, table := range tables {
		fields := getField(table.Name)
		generateModel(table, fields)
	}
}

// 生成model.go
func generateModel(table Table, fields []Field) {
	// 结构体内容
	content := "package model\n\n"
	// 表注释
	if len(table.Comment) > 0 {
		content += "// " + table.Comment + "\n"
	}
	// 结构体名称
	content += "type " + generator.CamelCase(table.Name) + " struct { \n"
	// 生成字段
	for _, field := range fields {
		fieldName := generator.CamelCase(field.Field) // 字段名
		fieldJson := `json:"` + field.Field + `"`
		fieldType := getFieldType(field)
		fieldComment := getFieldComment(field)
		content += "	" + fieldName + "  " + fieldType + "  `" + fieldJson + "`  " + fieldComment + "\n"
	}
	content += "}"
	// 文件名
	fileName := conf.ModelPath + generator.CamelCase(table.Name) + ".go"
	// 生成文件
	var file *os.File
	var err error
	if fileIsExist(fileName) { // 是否存在
		if !conf.ModelReplace { // 是否覆盖源文件
			fmt.Println(generator.CamelCase(table.Name) + " 已存在，需删除才能重新生成...")
			return
		}
		// 覆盖源文件
		file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_RDONLY, 0666)
		if err != nil {
			panic(err)
		}
	} else {
		file, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()
	// 写入文件
	_, err = io.WriteString(file, content)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(generator.CamelCase(table.Name) + "模特妹妹来了")
	}
}

// 检查文件是否存在
func fileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// 不存在
		return false
	}
	return true
}

// 获取字段注释
func getFieldComment(field Field) string {
	if len(field.Comment) > 0 {
		return "// " + field.Comment
	} else {
		return ""
	}
}

// 获取字段类型
func getFieldType(field Field) string {
	typeArr := strings.Split(field.Type, "(")
	switch typeArr[0] {
	case "int":
		return "int"
	case "integer":
		return "int"
	case "mediumint":
		return "int"
	case "bit":
		return "int"
	case "year":
		return "int"
	case "smallint":
		return "int"
	case "tinyint":
		return "int"
	case "bigint":
		return "int64"
	case "decimal":
		return "float32"
	case "double":
		return "float32"
	case "float":
		return "float32"
	case "real":
		return "float32"
	case "numeric":
		return "float32"
	case "timestamp":
		return "time.Time"
	case "datetime":
		return "time.Time"
	case "time":
		return "time.Time"
	default:
		return "string"
	}
}

// 获取表的所有字段
func getField(tableName string) []Field {
	// show full columns from auths
	var fields []Field
	tools.GetDB().Raw("show full columns from " + tableName).Find(&fields)
	return fields
}

// 获取当前数据库的所有表
func getTables(tableNamesStr string) []Table {
	var tables []Table
	if tableNamesStr == "" {
		tools.GetDB().
			Raw("select table_name as name ,table_comment as comment "+
				"from information_schema.TABLES where TABLE_SCHEMA = ?", conf.DbConf.Db).
			Find(&tables)
	} else {
		tools.GetDB().
			Raw("select table_name as name ,table_comment as comment " +
				"from information_schema.TABLES where TABLE_NAME in (" + tableNamesStr + ")").
			Find(&tables)
	}
	return tables
}
