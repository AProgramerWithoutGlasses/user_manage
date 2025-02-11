package dao

import "fmt"

func (dao *Dao) GetHeadUrl(username string) (headUrl string, err error) {
	stmt, err := dao.db.Prepare("select head_url from user where username=?")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&headUrl)
	if err != nil {
		fmt.Println("stmt.QueryRow() err", err)
		return
	}

	return
}

func (dao *Dao) UpdateHeadUrl(headUrl string, username string) (err error) {
	stmt, err := dao.db.Prepare("UPDATE user SET head_url = ? WHERE username = ?")
	if err != nil {
		fmt.Println("dao.db.Prepare() err", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(headUrl, username)
	if err != nil {
		fmt.Println("dao.db.Exec() err", err)
		return
	}

	return
}
