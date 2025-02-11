package server

import (
	"html/template"
	"net/http"
	"user_manage/src/logger"
)

func failServer(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/views/fail.html")
	if err != nil {
		logger.Error("template.ParseFiles() err", err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		logger.Error("template.Execute() err", err)
		return
	}

	return
}
