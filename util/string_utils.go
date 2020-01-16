package util

import (
	"regexp"
	"strings"
	"unicode"
)

// 首字母变小写
func LowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// 首字母变大写
func UpperFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 下划线转驼峰命名  xx_yy --> XxYy
func CamelString(str string) string {
	strArr := strings.Split(str, "_")
	camelStr := ""
	for _, s := range strArr {
		camelStr += UpperFirst(s)
	}
	return camelStr
}

// 下划线转驼峰命名强化版  xx_yy --> XxYy
func CamelStringReg(str string) string {
	re := regexp.MustCompile("(_|-)([a-zA-Z]+)")
	camel := re.ReplaceAllString(str, " $2")
	camel = strings.Title(camel)
	camel = strings.Replace(camel, " ", "", -1)

	return camel
}

// 驼峰转下划线 XxYxZz --> xx_yy_zz
func SnakeString(str string) string {
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snakeString := matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return snakeString
}
