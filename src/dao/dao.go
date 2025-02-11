package dao

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"user_manage/src/pkg/settings"
)

type Dao struct {
	db  *sql.DB
	rdb *redis.Client
}

// Init 用于初始化 db 和 rdb，并返回这二者
func Init(app *settings.AppConfig) *Dao {
	dao := &Dao{
		db:  initDB(app.MySQLConfig),
		rdb: initRDB(app.RedisConfig),
	}

	return dao
}
