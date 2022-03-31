package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
}

func Boot() {
	fmt.Println("log boot start")
	initConfig()
	fmt.Println("log boot finished")
}

func initConfig() *logrus.Logger {
	now := time.Now()
	logFilePath := viper.GetString("log.log_folder_path")
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
	//设置输出
	Log.Out = src
	//设置日志级别
	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	return Log
}
