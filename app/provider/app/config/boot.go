package config

import (
	"chenwlnote.gin-api/app/provider/app"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var config = new(Config)
var server *app.Server

type Config struct {
	HttpServer struct {
		Addr        string `json:"addr"`
		HttpPidFile string `json:"http_pid_file"`
	} `json:"http_server"`
	Database struct {
		Mysql struct {
			Read  []MysqlInfo `json:"read"`
			Write []MysqlInfo `json:"write"`
		} `json:"mysql"`
		Redis struct {
			Master []RedisInfo `json:"master"`
			Slave  []RedisInfo `json:"slave"`
		}
	} `json:"database"`
	Log struct {
		ApiLogFile string `json:"api_log_file"`
		ApiLogPath string `json:"api_log_path"`
	} `json:"log"`
	Trace TraceConfig `json:"trace"`
}

type RedisInfo struct {
	Host            string        `json:"host"`
	Port            int           `json:"port"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	Db              int           `json:"db"`
	IdleTimeout     time.Duration `json:"idle_timeout"`
	MaxConnLifetime time.Duration `json:"max_conn_lifetime"`
	MaxIdle         int           `json:"max_idle"`
	MaxActive       int           `json:"max_active"`
}

type MysqlInfo struct {
	Host            string           `json:"host"`
	Port            int              `json:"port"`
	Username        string           `json:"username"`
	Password        string           `json:"password"`
	DatabaseConfig  []DatabaseConfig `json:"database_config"`
	ConnMaxIdleTime time.Duration    `json:"conn_max_idle_time"`
	ConnMaxLifeTime time.Duration    `json:"conn_max_life_time"`
	MaxIdleConns    int              `json:"max_idle_conns"`
	MaxOpenConns    int              `json:"max_open_conns"`
}

type DatabaseConfig struct {
	Alias   string `json:"alias"`
	Name    string `json:"name"`
	Charset string `json:"charset"`
}

type TraceConfig struct {
	Driver     string `json:"driver"`
	AppName    string `json:"app_name"`
	ServerAddr string `json:"server_addr"`
	Open       bool   `json:"open"`
}

func init() {
	server = app.Get().Server()
	viper.SetConfigName("app_" + server.GetEnv())
	viper.SetConfigType("json")
	viper.AddConfigPath(server.GetProjectPath() + "/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper read config fail")
	}

	if readConfigErr := viper.ReadInConfig(); readConfigErr != nil {
		panic(readConfigErr)
	}

	initHttpServerConfig()
	initDatabaseConfig()
	initLogConfig()
	initTraceConfig()
	fmt.Println("config init finished")
}

func initTraceConfig() {
	traceConfigBytes, jsonErr := json.Marshal(viper.Get("log"))
	if jsonErr != nil {
		panic("解析 trace config 异常:" + jsonErr.Error())
	}
	errRead := json.Unmarshal(traceConfigBytes, &config.Trace)
	if errRead != nil {
		panic("解析 trace 异常" + errRead.Error())
	}
	fmt.Println("trace config init finished")
}

func initLogConfig() {
	logConfigBytes, jsonErr := json.Marshal(viper.Get("log"))
	if jsonErr != nil {
		panic("解析 log config 异常:" + jsonErr.Error())
	}
	errRead := json.Unmarshal(logConfigBytes, &config.Log)
	if errRead != nil {
		panic("解析 log 异常" + errRead.Error())
	}
	fmt.Println("log config init finished")
}

func initHttpServerConfig() {
	httpServerConfigBytes, jsonErr := json.Marshal(viper.Get("http_server"))
	if jsonErr != nil {
		panic("解析 http server config 异常:" + jsonErr.Error())
	}
	errRead := json.Unmarshal(httpServerConfigBytes, &config.HttpServer)
	if errRead != nil {
		panic("解析 http server 异常" + errRead.Error())
	}
	fmt.Println("http server config init finished")
}

func initDatabaseConfig() {
	databaseConfigBytes, jsonErr := json.Marshal(viper.Get("database"))
	if jsonErr != nil {
		panic("解析 config 异常:" + jsonErr.Error())
	}
	errRead := json.Unmarshal(databaseConfigBytes, &config.Database)
	if errRead != nil {
		panic("解析db 异常" + errRead.Error())
	}
	fmt.Println("database config init finished")
}

func Boot() {

}

func Get() Config {
	return *config
}
