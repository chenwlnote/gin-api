package logger

import "chenwlnote.gin-api/app/provider/app/log"

func Info(title string, a map[string]interface{}) {
	log.Log.WithFields(a).Info(title)
}
