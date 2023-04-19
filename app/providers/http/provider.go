package http

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

type Provider struct {
	app foundation.Application
	*foundation.UnimplementedProvider
}

var _ foundation.Provider = (*Provider)(nil)

func NewProvider(app foundation.Application) *Provider {
	return &Provider{
		app: app,
	}
}

func (g *Provider) Register() {
	g.app.Singleton("http.kernel", func(container container.Container) interface{} {
		return gin.Default()
	})
}

func (g *Provider) Boot() {
	g.router()
	g.middleware()
}

func (g *Provider) router() {
	for _, router := range routers {
		router.Register(facades.Http())
	}
}

func (g *Provider) middleware() {
	for _, middleware := range middlewares {
		facades.Http().Use(middleware)
	}
}
