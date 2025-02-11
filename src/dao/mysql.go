package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"user_manage/src/pkg/settings"
)

func initDB(m *settings.MySQLConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.DB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("gorm init failed %v", zap.Error(err))
	}

	// 检查连接
	if err = db.Ping(); err != nil {
		zap.L().Error("连接数据库失败:", zap.Error(err))
	}
	fmt.Println("数据库连接成功")

	return db
}
