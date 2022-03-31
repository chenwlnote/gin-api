package limit

import (
	"fun.tvapi/app/exception"
	"time"
)

func RunTime(f func() (interface{}, error), duration time.Duration) (data interface{}, err error) {
	var isOk = make(chan bool)
	var runErr error = nil
	var result interface{} = nil
	go func() {
		go func() {
			<-time.NewTimer(duration).C
			runErr = &exception.FuncTimeoutError{FunError: exception.FunError{Code: 1000}}
			isOk <- true
		}()
		var err error = nil
		result, err = f()
		if err != nil {
			runErr = err
		}
		isOk <- true
	}()
	<-isOk
	return result, runErr
}
