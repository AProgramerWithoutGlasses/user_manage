package dao

import (
	"database/sql"
	"fmt"
	"user_manage/src/model"
)

func (dao *Dao) GetPermissionByUsername(username string) (permission string, err error) {
	stmt, err := dao.db.Prepare("select permission from user where username=?")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&permission)
	if err != nil {
		fmt.Println("dao.Query() err", err)
		return
	}

	return
}

func (dao *Dao) InsertUser(registerModel *model.RegisterModel) (err error) {
	stmt, err := dao.db.Prepare("INSERT INTO user (username, password, name, age, gender, permission, head_url) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(registerModel.Username, registerModel.Password, registerModel.Name, registerModel.Age, registerModel.Gender, model.DefaultPermission, model.DefaultHeadUrl)
	if err != nil {
		fmt.Println("dao.db.Exec() err", err)
		return
	}

	return
}

// ExistedUser 检查User表中是否存在传入username对应的记录，若存在则返回true，不存在返回false
func (dao *Dao) ExistedUser(username string) (existed bool, err error) {
	stmt, err := dao.db.Prepare("select count(1) from user where username = ?")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close()

	var count int // 查到的记录数量
	err = stmt.QueryRow(username).Scan(&count)
	if err != nil {
		fmt.Println("dao.db.Query() err", err)
		return
	}

	if count > 0 {
		existed = true
	}
	return
}

func (dao *Dao) GetPwdByUsername(username string) (dbPassword string, err error) {
	stmt, err := dao.db.Prepare("select password from user where username=?")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&dbPassword)
	if err != nil {
		fmt.Println("dao.Query() err", err)
		return
	}

	return
}

func (dao *Dao) GetAllUser() (users []model.User, err error) {
	// 准备 SQL 查询语句
	stmt, err := dao.db.Prepare("SELECT username, password, name, age, gender, permission FROM user")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close()

	// 执行查询
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("dao.db.Query() err", err)
		return
	}
	defer rows.Close()

	// 遍历结果集
	for rows.Next() {
		var user model.User
		if err = rows.Scan(&user.Username, &user.Password, &user.Name, &user.Age, &user.Gender, &user.Permission); err != nil {
			fmt.Println("dao.db.Query() err", err)
			return
		}
		users = append(users, user) // 将用户添加到切片中
	}

	// 检查遍历过程中的错误
	if err = rows.Err(); err != nil {
		fmt.Println("rows.Err()", err)
		return
	}

	return
}

func (dao *Dao) EditUser(editModel model.EditModel) (result sql.Result, err error) {
	// 准备 SQL 更新语句
	stmt, err := dao.db.Prepare("UPDATE user SET username = ?, name = ?, age = ?, gender = ?, permission = ? WHERE username = ?")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close()

	// 执行更新
	result, err = stmt.Exec(editModel.Username, editModel.Name, editModel.Age, editModel.Gender, editModel.Permission, editModel.OriginUsername)
	if err != nil {
		fmt.Println("stmt.Exec() err", err)
		return
	}

	return
}

func (dao *Dao) DeleteUser(username string) (result sql.Result, err error) {
	// 准备删除语句
	stmt, err := dao.db.Prepare("DELETE FROM user WHERE username = ?")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close() // 确保在函数结束时关闭语句

	// 执行删除操作
	result, err = stmt.Exec(username)
	if err != nil {
		fmt.Println("dao.db.Exec() err", err)
		return
	}

	return
}
