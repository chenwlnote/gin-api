package router

import (
	v1 "fun.tvapi/router/v1"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) {
	v1.Load(g)
}
