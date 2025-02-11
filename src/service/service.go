package service

import (
	"golang.org/x/sync/singleflight"
	"user_manage/src/dao"
	"user_manage/src/pkg/settings"
)

type Service struct {
	dao    *dao.Dao
	single *singleflight.Group // 合并相同的并发请求, 提高性能
}

func InitService(app *settings.AppConfig) *Service {
	svc := &Service{
		dao:    dao.Init(app),
		single: new(singleflight.Group),
	}
	return svc

}
