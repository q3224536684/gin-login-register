package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var _db *gorm.DB
var once sync.Once

func InitDb() (*gorm.DB, error) {
	config := GetConfig()
	mysqlConf := config.MySql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Db, mysqlConf.Timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxOpenConns(100) // 设置数据库连接池最大连接数
	sqlDb.SetMaxIdleConns(20)  // 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return db, nil
}

func GetDb() *gorm.DB {
	once.Do(func() {
		db, err := InitDb()
		if err != nil {
			log.Fatal("连接数据库失败, error: ", err)
		}
		_db = db
	})
	return _db
}
