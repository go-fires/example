package controllers

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (c *IndexController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello World")
}
