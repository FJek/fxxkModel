package conf


// todo 配置表名


// 文件生成的路径
const ModelPath = "./model/"

// 是否覆盖已存在model
const ModelReplace = true

/**
数据库配置
*/
type BaseConf struct {
	Host string // 主机
	Port string // 端口
	User string // 用户名
	Pwd  string // 密码
	Db   string // 数据库名
}

/**
配置实例
*/
var DbConf BaseConf = BaseConf{
	Host: "localhost",
	Port: "3306",
	User: "root",
	Pwd:  "",
	Db:   "for-change",
}
