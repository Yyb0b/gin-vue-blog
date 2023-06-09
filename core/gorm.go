package core

import (
	"fmt"
	"gin_vue_project/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

//func Gorm() *gorm.DB {
//	return MysqlConnect()
//}

// gorm连接数据库
func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warn("未配置mysql, 取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqllogger logger.Interface
	if global.Config.System.Env == "dev" {
		//开发环境显示所有sql
		mysqllogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqllogger = logger.Default.LogMode(logger.Error) //只打印错误的sql
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqllogger,
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最大可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
