package config

import (
	"github.com/gofiber/storage/redis/v3"
	"runtime"
)

var Store *redis.Storage

func NewRedis() {
	Store = redis.New(redis.Config{
		Host:      "127.0.0.1",
		Port:      6379,
		Username:  "",
		Password:  "",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
}

func CloseRedis() error {
	return Store.Close()
}
