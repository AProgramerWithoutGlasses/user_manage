package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"user_manage/src/model"
)

// GetHead 从数据库中拿到文件的URL
func (s *Service) GetHead(username string) (headUrl string, err error) {
	headUrl, err = s.dao.GetHeadUrl(username)
	if err != nil {
		fmt.Println("GetHeadUrl err:", err)
	}

	//file, err = ioutil.ReadFile(headUrl)
	//if err != nil {
	//	fmt.Println("ReadFile err:", err)
	//	return
	//}

	return
}

func (s *Service) HeadService(file io.Reader, fileHeader *multipart.FileHeader, username string) (err error) {
	// 解析 filePath
	filePath := getFilePath(fileHeader)

	// 存入本地
	err = saveLocalFile(file, filePath)
	if err != nil {
		fmt.Println("SaveLocalFile err:", err)
		return
	}

	// 存入数据库
	filePath = "/" + filePath // 开头加上 '/'
	err = s.dao.UpdateHeadUrl(filePath, username)
	if err != nil {
		fmt.Println("dao.UpdateHeadUrl err: ", err)
		return
	}

	fmt.Println("文件写入成功: ", filePath)

	return
}

func saveLocalFile(file io.Reader, filePath string) (err error) {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ioutil.ReadAll err: ", err)
		return
	}

	// 写入文件
	err = ioutil.WriteFile(filePath, b, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil.WriteFile() err: ", err)
		return
	}

	return
}

func getFilePath(fileHeader *multipart.FileHeader) (filePath string) {
	// 使用相对路径
	filePath = filepath.Join(model.ReferenceDir, fileHeader.Filename)

	// 用 '/' 替换 '\'
	filePath = strings.Replace(filePath, "\\", "/", -1)

	return
}
