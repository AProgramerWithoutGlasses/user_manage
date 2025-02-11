package server

import (
	"html/template"
	"net/http"
	"user_manage/src/logger"
	"user_manage/src/pkg/response"
)

func headUpdateServer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/views/index.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		response.FailView(w, r)
		return
	}

	// 接收
	file, fileHeader, err := r.FormFile("avatar")
	if err != nil {
		logger.Error("r.FormFile(avatar) err: ", err)
		data, _ := getIndexData(r)
		response.Fail(w, t, err.Error(), data)
		return
	}
	defer file.Close()

	username, err := r.Cookie("username")
	if err != nil {
		response.Fail(w, t, err.Error(), nil)
		logger.Error("r.Cookie(username) err: ", err)
		return
	}

	// 业务
	err = svc.HeadService(file, fileHeader, username.Value)
	if err != nil {
		logger.Error("svc.HeadService(file, fileHeader, username.Value) err: ", err)
		data, _ := getIndexData(r)
		response.Fail(w, t, err.Error(), data)
		return
	}

	data, err := getIndexData(r)
	if err != nil {
		logger.Error("getIndexData() err", err)
		response.Fail(w, t, err.Error(), data)
		return
	}

	// 响应
	response.Success(w, t, "头像更换成功", data)
}
