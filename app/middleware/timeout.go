package middleware

import (
	"fmt"
	"fun.tvapi/app/pkg/util/limit"
	"fun.tvapi/app/response"
	"github.com/gin-gonic/gin"
	"time"
)

func Timeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			_, err := limit.RunTime(func() (interface{}, error) {
				c.Next()
				return c, nil
			}, timeout)
			if err != nil {
				fmt.Println("api timeout:", err)
				response.Timeout(c)
				c.Abort()
				return
			}
			c.Next()
		}()
	}
}
