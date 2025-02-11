package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"user_manage/src/pkg/settings"
)

var logger *logrus.Logger

func Init(cfg *settings.LogConfig) (err error) {
	// 创建一个新的 Logrus Logger
	logger = logrus.New()

	// 设置日志输出到文件
	file, err := os.OpenFile(cfg.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal(err)
	}

	logger.SetOutput(file)

	// 设置日志格式为 JSON
	logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	// 记录日志
	logger.Info("应用程序启动")
	fmt.Println("日志启动成功")

	return
}

func Info(a ...string) {
	fmt.Println(a)
	logger.Info(a)
}

func Error(a ...any) {
	fmt.Println(a)
	logger.Error(a)
}
