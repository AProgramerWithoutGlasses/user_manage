package err

import "user_manage/src/logger"

// PanicRecover 用于在程序panic报错停止后捕获panic
func PanicRecover() {
	if err := recover(); err != nil {
		logger.Error("panic recover: ", err)
	}
}
