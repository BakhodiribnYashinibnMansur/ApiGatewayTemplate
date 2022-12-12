package handler

import (
	"github.com/gin-gonic/gin"
)

func Ping(api *gin.Engine, handler *Handler) {
	api.GET("/", handler.Ping)
}

func (handler *Handler) Ping(ctx *gin.Context) {
	ctx.JSON(200, "SUCCESSFUL Pong")
}
