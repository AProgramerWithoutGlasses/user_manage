package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"user_manage/src/middleware"
	"user_manage/src/pkg/settings"
	"user_manage/src/service"
)

var svc *service.Service

func Init(app *settings.AppConfig, svc1 *service.Service) {
	svc = svc1 // 依赖注入

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/login", loginServer)
	r.HandleFunc("/logout", logoutServer)
	r.HandleFunc("/register", registerServer)
	r.HandleFunc("/index", middleware.MidOnlineStatus(indexServer))
	r.HandleFunc("/edit", editServer)
	r.HandleFunc("/delete", deleteServer)
	r.HandleFunc("/head_update", headUpdateServer)
	r.HandleFunc("/index/logmsg", logMsgServer)
	r.HandleFunc("/fail", failServer)

	http.ListenAndServe("localhost:"+strconv.Itoa(app.Port), r)
}

//r.Router("DELETE","users/{key}") // 删
//r.Router("GET","users")	// 查集合
//r.Router("GET","users/{key}")	// 查单个
//r.Router("POST","users")	// 新增
//r.Router("UPDATE","users/{key}")	// 编辑
