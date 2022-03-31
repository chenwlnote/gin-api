package database

import (
	"fun.tvapi/app/provider/app/config"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"strconv"
	"time"
)

var redisReadPool []*redis.Pool
var redisWritePool []*redis.Pool

type RedisPool struct {
}

func init() {
	initRedisReadInstance()
	initRedisWriteInstance()
	checkRedisPool()
}

func checkRedisPool() {
	if redisReadPool == nil || redisWritePool == nil {
		panic("redis pool 初始化失败")
	}
}

func initRedisReadInstance() {
	for _, redisConfig := range config.Get().Database.Redis.Slave {
		redisReadPool = append(redisReadPool, newPool(redisConfig))
	}
}

func newPool(redisConfig config.RedisInfo) *redis.Pool {
	return &redis.Pool{
		MaxIdle:         redisConfig.MaxIdle,
		IdleTimeout:     redisConfig.IdleTimeout * time.Second,
		MaxActive:       redisConfig.MaxActive,
		MaxConnLifetime: redisConfig.MaxConnLifetime,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisConfig.Host+":"+strconv.Itoa(redisConfig.Port))
		},
	}
}

func initRedisWriteInstance() {
	for _, redisConfig := range config.Get().Database.Redis.Master {
		redisWritePool = append(redisWritePool, newPool(redisConfig))
	}
}

func NewRedisPool() *RedisPool {
	return &RedisPool{}
}

func (pool *RedisPool) Read() *redis.Pool {
	if redisReadPool == nil || len(redisReadPool) == 0 {
		panic("redis read pool is nil")
	}
	if len(redisReadPool) == 1 {
		return redisReadPool[0]
	}
	index := rand.Intn(len(redisReadPool) - 1)
	return redisReadPool[index]
}

func (pool *RedisPool) Write() *redis.Pool {
	if redisWritePool == nil || len(redisWritePool) == 0 {
		panic("redis write pool is nil")
	}
	if len(redisWritePool) == 1 {
		return redisWritePool[0]
	}
	index := rand.Intn(len(redisWritePool) - 1)
	return redisWritePool[index]
}
