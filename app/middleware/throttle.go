package middleware

import (
	"fmt"
	"fun.tvapi/app/response"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

func Throttle(fillInterval time.Duration, capacity int64, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, capacity, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			response.Forbidden(c)
			fmt.Println("")
			c.Abort()
			return
		}
		c.Next()
	}
}
