package httpserver

import (
	"fmt"
	"fun.tvapi/app/middleware"
	"fun.tvapi/router"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	r *gin.Engine
)

func init() {
	fmt.Println("server boot init")
}

func Boot() {
	r = gin.Default()
	pprof.Register(r) // 性能分析
	gin.ForceConsoleColor()
	r.Use(middleware.Timeout(2 * time.Second))
	r.Use(middleware.CatchError())
	r.Use(gin.Recovery())
	r.Use(middleware.Throttle(1*time.Second, 1000, 10))
	r.Use(middleware.Logger())
	r.Use(middleware.RequestId())
	router.Load(r)
}
