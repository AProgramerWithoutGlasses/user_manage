package service

import (
	"fmt"
	"user_manage/src/model"
)

func (s *Service) RegisterService(registerModel *model.RegisterModel) (err error) {
	// 先判断该用户是否存在
	existed, err := s.dao.ExistedUser(registerModel.Username)
	if err != nil {
		fmt.Println("dao.ExistedUser err:", err)
		return
	}
	fmt.Println("dao.Existed is:", existed)

	if existed {
		// 若该用户已存在，则返回，不再注册。
		err = model.ErrUserExist
	} else {
		// 若该用户不存在，则进行注册
		err = s.dao.InsertUser(registerModel)
		if err != nil {
			fmt.Println("dao.InsertUser err:", err)
			return
		}
	}

	return
}
