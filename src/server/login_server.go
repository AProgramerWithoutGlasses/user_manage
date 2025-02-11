package server

import (
	"html/template"
	"net/http"
	"user_manage/src/logger"
	"user_manage/src/model"
	"user_manage/src/pkg/response"
	"user_manage/src/pkg/validate"
)

func loginServer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/views/login.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		response.FailView(w, r)
		return
	}

	switch r.Method {
	case "GET":
		err = t.Execute(w, nil)
		if err != nil {
			logger.Error("template.Execute() err", err)
			response.FailView(w, r)
			return
		}

	case "POST":
		// 接收参数
		var loginModel model.LoginModel
		loginModel.Username = r.FormValue("username")
		loginModel.Password = r.FormValue("password")

		// 验证参数
		err = validate.Validate(loginModel)
		if err != nil {
			logger.Error("validate.Validate() err", err)
			response.ParamUnvalid(w, t, err.Error(), "")
			return
		}

		// 业务
		myPermission, err := svc.LoginService(loginModel)
		if err != nil {
			logger.Error("svc.LoginService() err", err)
			response.Fail(w, t, err.Error(), nil)
			return
		}

		// 设置cookie
		c1 := http.Cookie{Name: "permission", Value: myPermission}
		c2 := http.Cookie{Name: "username", Value: loginModel.Username}
		c3 := http.Cookie{Name: "loginStatus", Value: model.OnlineStatus} // 登录状态 // 1
		http.SetCookie(w, &c1)
		http.SetCookie(w, &c2)
		http.SetCookie(w, &c3)

		// 响应
		response.ReDirect(w, r, "/index")
	}

	return
}

// 退出登录
func logoutServer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/views/login.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		response.FailView(w, r)
		return
	}

	c3 := http.Cookie{Name: "loginStatus", Value: model.OfflineStatus} // 登录状态
	http.SetCookie(w, &c3)

	err = t.Execute(w, nil)
	if err != nil {
		logger.Error("template.Execute() err", err)
		response.FailView(w, r)
		return
	}
}
