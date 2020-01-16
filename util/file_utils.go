package util

import "os"

// 检查文件是否存在
func FileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// 不存在
		return false
	}
	return true
}
