package model

import "errors"

var (
	ErrUserNotExist    = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("密码错误")
	ErrUserExist       = errors.New("用户存在")
)

var (
	DefaultPermission = "user"    // 默认权限-用户
	OfflineStatus     = "online"  // 登录状态-在线
	OnlineStatus      = "offline" // 登录状态-离线
)

type LoginModel struct {
	Username string `json:"username" validate:"required,no-whitespace"`
	Password string `json:"password" validate:"required,no-whitespace"`
}

type RegisterModel struct {
	Username string `json:"username" validate:"required,numeric"`  // numeric 必须是数字
	Password string `json:"password" validate:"required,alphanum"` // alphanum 必须是数字字母组合
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required,gte=0,lte=100,numeric"`
	Gender   string `json:"gender" validate:"required,oneof=男 女"`
}

type User struct {
	Username   string `json:"username" validate:"required,numeric"`  // numeric 必须是数字
	Password   string `json:"password" validate:"required,alphanum"` // alphanum 必须是数字字母组合
	Name       string `json:"name" validate:"required"`
	Age        string `json:"age" validate:"required,gte=0,lte=100,numeric"`
	Gender     string `json:"gender" validate:"required,oneof=男 女"`
	Permission string `json:"permission" validate:"required,oneof=user manager"`
}

type IndexModel struct {
	Users        []User `json:"users"`
	MyPermission string `json:"my_permission"`
	MyHead       string `json:"myHead"`
	MyUsername   string `json:"my_username"`
}

type EditModel struct {
	OriginUsername string `json:"originUsername" validate:"required,numeric"`
	Username       string `json:"username" validate:"required,numeric"` // numeric 必须是数字
	Name           string `json:"name" validate:"required"`
	Age            int    `json:"age" validate:"required,gte=0,lte=100,numeric"`
	Gender         string `json:"gender" validate:"required,oneof=男 女"`
	Permission     string `json:"permission" validate:"required,oneof=user manager"`
}

type DeleteModel struct {
	Username string `json:"username" validate:"required,numeric"`
}
