package service

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"user_manage/src/model"
)

func (s *Service) LogMsgService(page int) (entries []model.LogEntry, totalEntries int, err error) {
	// 打开文件
	file, err := os.Open("app.log")
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 创建一个 Scanner 逐行读取
	scanner := bufio.NewScanner(file)
	entries = make([]model.LogEntry, 0)

	// 读取所有日志条目
	for scanner.Scan() {
		var entry model.LogEntry
		line := scanner.Text()
		if err := json.Unmarshal([]byte(line), &entry); err == nil {
			entries = append(entries, entry)
		} else {
			log.Printf("解析错误: %v", err)
		}
	}

	// 检查扫描中的错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("读取文件时出错: %v", err)
	}

	// 设置总条目数
	totalEntries = len(entries)

	// 计算分页
	start := (page - 1) * 10
	end := start + 10
	if end > totalEntries {
		end = totalEntries
	}
	if start > totalEntries {
		start = totalEntries
	}

	return entries[start:end], totalEntries, nil
}
