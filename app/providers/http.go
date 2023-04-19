package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-fires/example/app/contracts/http"
	"github.com/go-fires/example/app/facades"
	"github.com/go-fires/example/app/http/routes"
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/contracts/foundation"
)

var routers = []http.Router{
	new(routes.Api),
}

var middlewares = []gin.HandlerFunc{
	gin.Logger(),
	gin.Recovery(),
}

type HttpProvider struct {
	app foundation.Application
	*foundation.UnimplementedProvider
}

func NewHttpProvider(app foundation.Application) *HttpProvider {
	return &HttpProvider{
		app: app,
	}
}

func (g *HttpProvider) Register() {
	g.app.Singleton("http.kernel", func(container container.Container) interface{} {
		return gin.Default()
	})
}

func (g *HttpProvider) Boot() {
	g.router()
	g.middleware()
}

func (g *HttpProvider) router() {
	for _, router := range routers {
		router.Register(facades.Http())
	}
}

func (g *HttpProvider) middleware() {
	for _, middleware := range middlewares {
		facades.Http().Use(middleware)
	}
}
