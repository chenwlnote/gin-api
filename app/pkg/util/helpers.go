package util

import uuid "github.com/satori/go.uuid"

func GenUUID() string {
	return uuid.NewV4().String()
}

func If(condition bool, trueVal interface{}, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
