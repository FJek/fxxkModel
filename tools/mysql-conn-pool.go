package tools

import (
	"fzw/fxxkModel/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

/**
数据连接操作
*/
type MysqlConnPool struct{}

var mysqlInstance *MysqlConnPool

/*
	获取连接的实例
	单例模式
*/
var mysqlOnce sync.Once

func GetInstance() *MysqlConnPool {
	mysqlOnce.Do(func() {
		mysqlInstance = &MysqlConnPool{}
	})
	return mysqlInstance
}

/*
	初始化数据库连接
*/
var db *gorm.DB
var dbErr error

func (*MysqlConnPool) InitMysqlConnPool() bool {
	//root:123456@/demo?charset=utf8mb4&parseTime=True&loc=Local"
	dbConf := conf.DbConf
	db, dbErr = gorm.Open(`mysql`, dbConf.User+`:`+dbConf.Pwd+`@/`+dbConf.Db+`?charset=utf8mb4&parseTime=True&loc=Local`)
	db.SingularTable(true)
	if dbErr != nil {
		log.Fatal(dbErr)
		return false
	}
	return true
}

func (*MysqlConnPool) GetMysqlDB() *gorm.DB {
	return db
}

// 对外方法获取 db
func GetDB() *gorm.DB {
	return GetInstance().GetMysqlDB()
}
