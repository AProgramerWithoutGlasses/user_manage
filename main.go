package main

import (
	"fmt"
	"user_manage/src/logger"
	err2 "user_manage/src/pkg/err"
	"user_manage/src/pkg/settings"
	"user_manage/src/pkg/validate"
	"user_manage/src/server"
	"user_manage/src/service"
)

func main() {
	// 全局 panic 捕获
	defer err2.PanicRecover()

	app, err := settings.Init()
	if err != nil {
		fmt.Printf("init setting failed,error: %v\n", err)
		return
	}

	err = logger.Init(app.LogConfig)
	if err != nil {
		fmt.Printf("init logger failed,error: %v\n", err)
		return
	}

	err = validate.Init()
	if err != nil {
		fmt.Printf("init validator failed,error: %v\n", err)
		return
	}

	// 依赖注入
	svc := service.InitService(app)
	server.Init(app, svc)

}
