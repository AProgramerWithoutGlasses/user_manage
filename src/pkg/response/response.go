package response

import (
	"fmt"
	"html/template"
	"net/http"
	"user_manage/src/logger"
)

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`  // 弹窗提示信息
	Data any    `json:"data"` // 渲染数据
}

func Success(w http.ResponseWriter, t *template.Template, msg string, data any) {
	message := &response{
		Code: 200,
		Msg:  msg,
		Data: data,
	}

	err := t.Execute(w, message)
	if err != nil {
		fmt.Println("response.Success() t.Execute() err: ", err)
		return
	}

	logger.Info(fmt.Sprintf("response message-->%+v", message))

	return
}

func Fail(w http.ResponseWriter, t *template.Template, msg string, data any) {
	message := &response{
		Code: 500,
		Msg:  msg,
		Data: data,
	}

	err := t.Execute(w, message)
	if err != nil {
		fmt.Println("response.Fail() t.Execute() err: ", err)
		return
	}

	logger.Info(fmt.Sprintf("response message-->%+v", message))

	return
}

func ParamUnvalid(w http.ResponseWriter, t *template.Template, msg string, data any) {
	message := &response{
		Code: 400,
		Msg:  msg,
		Data: data,
	}

	err := t.Execute(w, message)
	if err != nil {
		fmt.Println("response.ParamUnvalid() t.Execute() err: ", err)
		return
	}

	logger.Info(fmt.Sprintf("response message-->%+v", message))

	return
}

func ReDirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func FailView(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/fail", http.StatusSeeOther)
}
