package service

import (
	"errors"
	"fmt"
	"user_manage/src/model"
)

func (s *Service) IndexService() (users []model.User, err error) {
	users, err = s.dao.GetAllUser()
	if err != nil {
		fmt.Println("get users err:", err)
		return
	}
	return
}

func (s *Service) EditService(editModel model.EditModel) (err error) {
	if editModel.Username != editModel.OriginUsername {
		existedUser, err := s.dao.ExistedUser(editModel.Username)
		if err != nil {
			fmt.Println("edit user error:", err)
			return err
		}

		if existedUser {
			err = errors.New("修改失败，该账号已存在")
			return err
		}
	}

	_, err = s.dao.EditUser(editModel)
	if err != nil {
		fmt.Println("s.dao.EditUser() err:", err)
		return
	}

	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	fmt.Println("result.RowsAffected() err", err)
	//	return
	//}
	//
	//if rowsAffected == 1 {
	//	fmt.Println("更新成功")
	//} else if rowsAffected == 0 {
	//	fmt.Println("更新异常，更新了0行记录")
	//	err = errors.New("更新异常，更新了0行记录")
	//} else {
	//	fmt.Println("更新异常，更新了超过1行记录")
	//	err = errors.New("更新异常，更新了超过1行记录")
	//}

	return
}

func (s *Service) DeleteService(username string, myUsername string) (err error) {
	// 检查该用户是否存在
	existed, err := s.dao.ExistedUser(username)
	if err != nil {
		fmt.Println("dao.ExistedUser() err", err)
		return
	}

	// 	检查该用户是否是自己
	if username == myUsername {
		err = errors.New("不能删除自己")
		return
	}

	// 若用户存在，则直接返回
	if !existed {
		err = errors.New("您想要删除的用户不存在")
		return
	}

	result, err := s.dao.DeleteUser(username)
	if err != nil {
		fmt.Println("s.dao.DeleteUser():", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("result.RowsAffected() err:", err)
		return
	}

	if rowsAffected == 1 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除异常，更新了超过1行记录")
		err = errors.New("删除异常，更新了超过1行记录")
	}

	return
}
