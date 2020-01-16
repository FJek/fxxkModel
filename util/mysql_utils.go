package util

import (
	"fmt"
	"fzw/fxxkModel/boot"
	"fzw/fxxkModel/conf"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"io"
	"os"
	"strings"
)


type Field struct {
	Field      string `gorm:"column:Field"`
	Type       string `gorm:"column:Type"`
	Null       string `gorm:"column:Null"`
	Key        string `gorm:"column:Key"`
	Default    string `gorm:"column:Default"`
	Extra      string `gorm:"column:Extra"`
	Privileges string `gorm:"column:Privileges"`
	Comment    string `gorm:"column:Comment"`
}

type Table struct {
	Name    string `gorm:"column:name"`
	Comment string `gorm:"column:comment"`
}

// 获取字段类型
func GetFieldType(field Field) string {
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
func GetField(tableName string) []Field {
	// show full columns from auths
	var fields []Field
	boot.GetDB().Raw("show full columns from " + tableName+"s").Find(&fields)
	return fields
}

// 获取当前数据库的所有表
func GetTables(tableNamesStr string) []Table {
	var tables []Table
	if tableNamesStr == "" {
		boot.GetDB().
			Raw("select table_name as name ,table_comment as comment "+
				"from information_schema.TABLES where TABLE_SCHEMA = ?", conf.DbConf.Db).
			Find(&tables)
	} else {
		boot.GetDB().
			Raw("select table_name as name ,table_comment as comment " +
				"from information_schema.TABLES where TABLE_NAME in (" + tableNamesStr + ")").
			Find(&tables)
	}
	// 去掉 表尾的's'
	for i,table := range tables {
		tables[i].Name = table.Name[:len(table.Name)-1]
	}
	return tables
}

// 获取字段注释
func getFieldComment(field Field) string {
	if len(field.Comment) > 0 {
		return "// " + field.Comment
	} else {
		return ""
	}
}

// 生成model.go
func GenerateModel(table Table, fields []Field) {
	// 结构体内容
	content := "package model\n\n"
	// 表注释
	if len(table.Comment) > 0 {
		content += "// " + table.Comment + "\n"
	}
	// 结构体名称
	content += "type " + CamelString(table.Name) + " struct { \n"
	// 生成字段
	for _, field := range fields {
		fieldName := generator.CamelCase(field.Field) // 字段名
		fieldJson := `json:"` + field.Field + `"`
		fieldType := GetFieldType(field)
		fieldComment := getFieldComment(field)
		content += "	" + fieldName + "  " + fieldType + "  `" + fieldJson + "`  " + fieldComment + "\n"
	}
	content += "}"
	// 文件名
	fileName := conf.ModelPath + LowerFirst(generator.CamelCase(table.Name)) + ".go"
	// 生成文件
	var file *os.File
	var err error
	if FileIsExist(fileName) { // 是否存在
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