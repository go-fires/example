package config

import (
	"github.com/go-fires/framework/redis"
	redis2 "github.com/redis/go-redis/v9"
)

var Redis = &redis.Config{
	Default: "default",
	Connections: map[string]redis.Configable{
		"default": &redis2.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	},
}
