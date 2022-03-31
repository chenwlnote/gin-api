package middleware

import (
	"fun.tvapi/app/provider/app/log"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := log.Log
	return func(c *gin.Context) {

		startTime := time.Now()

		c.Next()

		latencyTime := time.Since(startTime)
		reqMethod := c.Request.Method
		urlPath := c.Request.URL.Path
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			urlPath,
		)
	}
}
