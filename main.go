package main

import "fzw/fxxkModel/tools"

import "fzw/fxxkModel/generate"

func main() {
	// 初始化数据库
	tools.Init()
	// 生成表信息
	//todo 表明可配置化
	generate.Generate("auths", "menus")
}
