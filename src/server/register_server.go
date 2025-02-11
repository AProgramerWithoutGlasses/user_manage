package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"user_manage/src/logger"
	"user_manage/src/model"
	"user_manage/src/pkg/response"
	"user_manage/src/pkg/validate"
)

func registerServer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/views/register.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		response.FailView(w, r)
		return
	}

	switch r.Method {
	case "GET":
		// 解析页面
		err = t.Execute(w, nil)
		if err != nil {
			logger.Error("template.Execute() err", err)
			response.FailView(w, r)
			return
		}

	case "POST":
		// 接收参数
		var registerModel model.RegisterModel
		registerModel.Username = r.FormValue("username")
		registerModel.Password = r.FormValue("password")
		registerModel.Name = r.FormValue("name")
		registerModel.Age, _ = strconv.Atoi(r.FormValue("age"))
		registerModel.Gender = r.FormValue("gender")

		logger.Info(fmt.Sprintf("%v\n", registerModel))

		// 参数校验
		err = validate.Validate(registerModel)
		if err != nil {
			logger.Error("validate.Validate() err", err)
			response.ParamUnvalid(w, t, err.Error(), "")
			return
		}

		// 业务
		err = svc.RegisterService(&registerModel)
		if err != nil {
			logger.Error("svc.RegisterService() err", err)
			response.Fail(w, t, err.Error(), nil)
			return
		}

		// 响应
		response.Success(w, t, registerModel.Username+"注册成功", nil)

		return
	}

}
