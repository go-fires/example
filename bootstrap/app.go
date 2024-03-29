package bootstrap

import (
	recoveryCustom "github.com/go-fires/example/app/debug/recovery"
	"github.com/go-fires/example/app/providers/http"
	"github.com/go-fires/example/config"
	"github.com/go-fires/framework/cache"
	"github.com/go-fires/framework/debug/recovery"
	"github.com/go-fires/framework/encryption"
	"github.com/go-fires/framework/foundation"
	"github.com/go-fires/framework/hashing"
	"github.com/go-fires/framework/redis"
)

func App() *foundation.Application {
	app := foundation.NewApplication()

	// Configure application
	app.Configure("http", config.Http)
	app.Configure("redis", config.Redis)
	app.Configure("cache", config.Cache)

	// Register providers
	app.Register(recovery.NewProvider(app, recovery.WithHandler(&recoveryCustom.Handler{})))
	app.Register(hashing.NewProvider(app))
	app.Register(redis.NewProvider(app))
	app.Register(encryption.NewProvider(app))
	app.Register(cache.NewProvider(app))
	app.Register(http.NewProvider(app))

	return app
}
