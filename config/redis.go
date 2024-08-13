package config

import (
	"github.com/gofiber/storage/redis/v3"
	"github.com/spf13/viper"
	"runtime"
)

func NewRedis(viper *viper.Viper) *redis.Storage {
	host := viper.GetString("redis.host")
	port := viper.GetInt("redis.port")
	password := viper.GetString("redis.password")
	username := viper.GetString("redis.username")

	Store := redis.New(redis.Config{
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  password,
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})

	return Store
}
