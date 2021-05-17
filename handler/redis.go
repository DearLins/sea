package handler

import (
	"github.com/go-redis/redis"
	"sea/conf"
)

var rdb *redis.Client

func InitClient() *redis.Client {
	config := conf.GetConfiguration()
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisHost+":"+config.RedisPort,
		Password: config.RedisPwd, // no password set
		DB:       config.RedisDatabase,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
