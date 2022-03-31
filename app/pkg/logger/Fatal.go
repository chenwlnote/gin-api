package logger

import "chenwlnote.gin-api/app/provider/app/log"

func Fatal(title string, a map[string]interface{}) {
	log.Log.WithFields(a).Fatal(title)
}
