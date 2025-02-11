package service

import (
	"database/sql"
	"errors"
	"fmt"
	"user_manage/src/model"
)

func (s *Service) LoginService(loginModel model.LoginModel) (myPermission string, err error) {
	dbPassword, err := s.dao.GetPwdByUsername(loginModel.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = model.ErrUserNotExist
		}
		fmt.Println("dao.db.QueryRow() err", err)
		return
	}

	if loginModel.Password != dbPassword {
		err = model.ErrInvalidPassword
	}

	myPermission, err = s.dao.GetPermissionByUsername(loginModel.Username)
	if err != nil {
		fmt.Println("dao.GetPermissionByUsername() err", err)
		return
	}

	return
}
