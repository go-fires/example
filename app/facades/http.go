package facades

import (
	"github.com/gin-gonic/gin"
	"github.com/go-fires/framework/facade"
)

func Http() *gin.Engine {
	var engine *gin.Engine

	if err := facade.App().Make("http.kernel", &engine); err != nil {
		panic(err)
	}

	return engine
}
