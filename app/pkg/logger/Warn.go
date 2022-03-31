package logger

import "chenwlnote.gin-api/app/provider/app/log"

func Warn(title string, a map[string]interface{}) {
	log.Log.WithFields(a).Warn(title)
}
