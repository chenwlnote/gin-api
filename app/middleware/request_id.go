package middleware

import (
	"chenwlnote.gin-api/app/pkg/util"
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = util.GenUUID()
		}
		ctx.Set("X-Request-Id", requestId)
		ctx.Writer.Header().Set("X-Request-Id", requestId)
		ctx.Next()
	}
}
