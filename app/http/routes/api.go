package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-fires/example/app/contracts/http"
	"github.com/go-fires/example/app/http/controllers"
)

type Api struct{}

var _ http.Router = (*Api)(nil)

func (a *Api) Register(e *gin.Engine) {
	index := new(controllers.IndexController)

	e.GET("/", index.Index)
}
