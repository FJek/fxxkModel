package main

import "fzw/fxxkModel/boot"

import "fzw/fxxkModel/generate"

func main() {
	// 初始化数据库
	boot.Init()
	// 生成表信息
	generate.Generate()
}
