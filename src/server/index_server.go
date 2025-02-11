package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"user_manage/src/logger"
	"user_manage/src/model"
	"user_manage/src/pkg/response"
	"user_manage/src/pkg/validate"
)

func indexServer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/views/index.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		response.FailView(w, r)
		return
	}

	// 业务
	data, err := getIndexData(r)
	if err != nil {
		logger.Error("getIndexData() err", err)
		response.Fail(w, t, err.Error(), data)
		return
	}

	// 响应
	response.Success(w, t, "", data)

	return
}

func editServer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/views/index.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		response.FailView(w, r)
		return
	}

	// 接收
	var editModel model.EditModel
	err = json.NewDecoder(r.Body).Decode(&editModel)
	if err != nil {
		logger.Error("json.NewDecoder(r.Body).Decode(&editModel)", err)
		data, _ := getIndexData(r)
		response.Fail(w, t, err.Error(), data)
		return
	}
	fmt.Printf("%+v\n", editModel)

	// 校验
	err = validate.Validate(editModel)
	if err != nil {
		logger.Error("validate.Validation() err", err)
		data, _ := getIndexData(r)
		response.ParamUnvalid(w, t, err.Error(), data)
		return
	}

	// 业务
	err = svc.EditService(editModel)
	if err != nil {
		logger.Error("svc.EditService() err", err)
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
	response.Success(w, t, "修改成功", data)

	return
}

func deleteServer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/views/index.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		response.FailView(w, r)
		return
	}

	// 接收
	var deleteModel model.DeleteModel
	err = json.NewDecoder(r.Body).Decode(&deleteModel)
	if err != nil {
		logger.Error("json.NewDecoder(r.Body).Decode(&deleteModel)", err)
		data, _ := getIndexData(r)
		response.Fail(w, t, err.Error(), data)
		return
	}
	fmt.Printf("deleteModel-->%+v\n", deleteModel)

	myUsername, err := r.Cookie("username")
	if err != nil {
		logger.Error("r.Cookie() err", err)
		return
	}

	// 校验
	err = validate.Validate(deleteModel)
	if err != nil {
		logger.Error("validate.Validation() err", err)
		data, _ := getIndexData(r)
		response.ParamUnvalid(w, t, err.Error(), data)
		return
	}

	// 业务
	err = svc.DeleteService(deleteModel.Username, myUsername.Value)
	if err != nil {
		logger.Error("svc.EditService() err", err)
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
	response.Success(w, t, "操作成功", data)

	return
}

func getIndexData(r *http.Request) (data model.IndexModel, err error) {
	// 业务 获取所有用户列表
	users, err := svc.IndexService()
	if err != nil {
		logger.Error("svc.IndexService() err", err)
		return
	}

	// 获取cookie
	permission, err := r.Cookie("permission")
	if err != nil {
		logger.Error("r.Cookie() err", err)
		return
	}

	username, err := r.Cookie("username")
	if err != nil {
		logger.Error("r.Cookie() err", err)
		return
	}

	loginStatus, err := r.Cookie("loginStatus")
	if err != nil {
		logger.Error("r.Cookie() err", err)
		return
	}

	// 检查登陆状态
	if loginStatus.Value == model.OfflineStatus {
		return model.IndexModel{}, errors.New("请您先登录")
	}

	// 获取myHeadImage
	myHead, err := svc.GetHead(username.Value)
	if err != nil {
		logger.Error("svc.GetHeadFile() err", err)
		return
	}

	data.Users = users
	data.MyPermission = permission.Value
	data.MyHead = myHead
	data.MyUsername = username.Value

	return
}
