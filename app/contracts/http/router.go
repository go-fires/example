package http

import "github.com/gin-gonic/gin"

type Router interface {
	Register(e *gin.Engine)
}
