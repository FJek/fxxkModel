package generate

import (
	"fzw/fxxkModel/conf"
	"fzw/fxxkModel/util"
)

/*
	生成model结构体
*/
func Generate() {
	tableNames := conf.TableNames.Names
	// 生成表信息
	tableNamesStr := ""
	for _, name := range tableNames {
		if tableNamesStr != "" {
			tableNamesStr += ","
		}
		tableNamesStr += "'" + name + "'"
	}
	tables := util.GetTables(tableNamesStr)
	// 生成模型
	for _, table := range tables {
		fields := util.GetField(table.Name)
		util.GenerateModel(table, fields)
	}
}
