package v1

import (
	DemoToolController "chenwlnote.gin-api/app/controller/tool/demo"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) {
	group := g.Group("v1")
	group.GET("/demo", DemoToolController.Index)
}
