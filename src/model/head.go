package model

var ReferenceDir = "static/img/"

var DefaultHeadUrl = "/static/img/封面.png"

type LogEntry struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Caller  string `json:"caller"`
	Msg     string `json:"msg"`
	Errors  string `json:"errors,omitempty"`
	Request string `json:"request,omitempty"`
	Stack   string `json:"stack,omitempty"`
}
