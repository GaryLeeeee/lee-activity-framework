package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var _db *gorm.DB
var _rd *redis.Client

func InitMySQL() {
	// MySQL连接参数
	username := "root"
	password := "88888888"
	host := "127.0.0.1"
	port := 3306
	dbName := "activity"
	timeout := "1s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbName, timeout)
	var err error
	_db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接MySQL失败，error=" + err.Error())
	}
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})
	_rd = client
}
