package exception

import "fmt"

const (
	ExecTimeOutException = 1000
)

var (
	errorCodeText = map[int]string{
		ExecTimeOutException: "执行超时",
	}
)

type FunError struct {
	Code int
}

func (err *FunError) Error() string {
	errMsg, ok := errorCodeText[err.Code]
	if !ok {
		errMsg = "系统异常"
	}
	return fmt.Sprintf("status code %d: %s", err.Code, errMsg)
}
