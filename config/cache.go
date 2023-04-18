package config

import (
	"github.com/go-fires/framework/cache"
	cacheContract "github.com/go-fires/framework/contracts/cache"
)

var Cache = &cache.Config{
	Default: "default",

	Stores: map[string]cacheContract.StoreConfigable{
		"default": &cache.MemoryStoreConfig{},
		"redis": &cache.RedisStoreConfig{
			Connection: "default",
			Prefix:     "cache",
		},
	},
}
