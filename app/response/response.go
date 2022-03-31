package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	RetCode int         `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{RetCode: 200, RetMsg: "success", Data: data})
}

func Error(c *gin.Context, err error) {
	c.JSON(400, Response{RetCode: -1, RetMsg: err.Error()})
}

func Exception(c *gin.Context, message string) {
	c.JSON(500, Response{RetCode: -1, RetMsg: message})
}

func Timeout(c *gin.Context) {
	c.JSON(502, Response{RetCode: -1, RetMsg: "网络异常"})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{RetCode: -1, RetMsg: "rate limit"})
}
