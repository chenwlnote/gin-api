package demo

import (
	"encoding/json"
	"fun.tvapi/app/dao/repository"
	"fun.tvapi/app/pkg/database"
	"fun.tvapi/app/response"
	"github.com/gin-gonic/gin"
	"time"
)

func Index(c *gin.Context) {
	r := repository.MediaRepository{}
	result := r.GetByIds([]int{1009017, 1009431, 1009445}, []string{"*"})
	conn := database.NewRedisPool().Write().Get()
	defer conn.Close()
	cache, _ := json.Marshal(result)
	conn.Do("set", "test_key", cache)
	response.Success(c, map[string]interface{}{"time": time.Now().Format("2006-01-02 15:04:05"), "result": result})
	return
}