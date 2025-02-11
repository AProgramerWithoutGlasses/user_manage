package validate

var errMsg = map[string]string{
	"no-whitespace": "不能含有空格",
	"required":      "不能为空",
	"numeric":       "必须是数字",
	"alphanumeric":  "只能包含字母和数字",
	"oneof":         "错误",
	"lte":           "超出限定范围",
	"gte":           "超出限定范围",
}

var fieldMsg = map[string]string{
	"Username":   "账号",
	"Password":   "密码",
	"Name":       "姓名",
	"Age":        "年龄",
	"Gender":     "性别",
	"Permission": "权限",
}
