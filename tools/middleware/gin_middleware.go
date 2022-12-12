package middleware

import (
	"github.com/gin-gonic/gin"
)


func GinMiddleware(route *gin.Engine) {
	route.Use(
		gin.Logger(),
		gin.Recovery(),
		gin.ErrorLogger(),
		CORSMiddleware(),
	)
}
