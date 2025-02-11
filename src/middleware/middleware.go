package middleware

import (
	"net/http"
	"user_manage/src/logger"
	"user_manage/src/model"
	"user_manage/src/pkg/response"
)

// MidOnlineStatus，用来记录每个请求，写进日子
func MidOnlineStatus(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 日志中间件
		loginStatus, err := r.Cookie("loginStatus")
		if err != nil {
			logger.Error("r.Cookie() err", err)
			return
		}

		// 检查登陆状态
		if loginStatus.Value == model.OfflineStatus {
			response.FailView(w, r)
		}

		handler(w, r)
	}
}
