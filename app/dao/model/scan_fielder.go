package model

type ScanFielder interface {
	ToScanField([]string) []interface{}
}
