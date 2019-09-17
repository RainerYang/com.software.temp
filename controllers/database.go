package controllers

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"com.motion.wallet/utils"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" //mysql enging
)

var (
	//DB mysql
	DB *gorm.DB
	//RedisW redis写
	RedisW *redis.Client
	//RedisR redis读
	RedisR *redis.Client
)

//RunMode debug
var RunMode = "debug"

func initDatabase() {
	var err error
	//init mysql client
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		GetConfigValue("mysqluser"),
		GetConfigValue("mysqlpassword"),
		GetConfigValue("mysqlhost"),
		utils.ParseInt(GetConfigValue("mysqlport")),
		GetConfigValue("mysqldatabase"),
		GetConfigValue("mysqlcharset"))

	DB, err = gorm.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxIdleConns(5)
	DB.DB().SetMaxOpenConns(20)
	err = DB.DB().Ping()
	if err != nil {
		panic(err)
	}

	RedisW = redis.NewClient(&redis.Options{
		Addr:         GetConfigValue("redisaddrwrite"),
		Password:     GetConfigValue("redispasswordwrite"),
		MaxRetries:   5,
		PoolSize:     80,
		MinIdleConns: 20,
		DB:           0,
	})

	RedisR = redis.NewClient(&redis.Options{
		Addr:         GetConfigValue("redisaddrread"),
		Password:     GetConfigValue("redispasswordread"),
		MaxRetries:   5,
		PoolSize:     80,
		MinIdleConns: 20,
		DB:           0,
	})
}
