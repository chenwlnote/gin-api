package middleware

import (
	"chenwlnote.gin-api/app/exception"
	"chenwlnote.gin-api/app/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				url := c.Request.URL
				method := c.Request.Method
				log.Printf("| url [%s] | method | [%s] | error [%s] |", url, method, err)
				var er exception.FunError
				err := json.Unmarshal([]byte(string(err.(string))), &er)
				if err != nil {
					c.JSON(400, response.Response{RetCode: er.Code, RetMsg: er.Error()})
					c.Abort()
					return
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
