package logger

import "fun.tvapi/app/provider/app/log"

func Warn(title string, a map[string]interface{}) {
	log.Log.WithFields(a).Warn(title)
}
