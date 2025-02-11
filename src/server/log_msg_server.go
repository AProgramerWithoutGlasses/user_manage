package server

import (
	"fmt"
	"html/template"
	"net/http"
	"user_manage/src/logger"
	"user_manage/src/model"
)

func logMsgServer(w http.ResponseWriter, r *http.Request) {
	// 先定义自定义函数
	funcMap := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
		"div": func(a, b int) int {
			return (a + b - 1) / b // 向上取整
		},
	}

	// 创建模板并注册自定义函数
	t, err := template.New("logmsg.html").Funcs(funcMap).ParseFiles("static/views/logmsg.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		return
	}

	// 获取当前页码，默认值为1
	page := 1
	if r.URL.Query().Get("page") != "" {
		fmt.Sscanf(r.URL.Query().Get("page"), "%d", &page)
	}

	// 业务逻辑
	entries, totalEntries, err := svc.LogMsgService(page)
	if err != nil {
		http.Error(w, "无法获取数据", http.StatusInternalServerError)
		return
	}

	// 响应数据
	data := struct {
		Entries      []model.LogEntry
		CurrentPage  int
		TotalEntries int
		TotalPages   int
	}{
		Entries:      entries,
		CurrentPage:  page,
		TotalEntries: totalEntries,
		TotalPages:   (totalEntries + 9) / 10, // 每页10条数据
	}

	t.Execute(w, data)
}
