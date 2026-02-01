package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/codex/growth-tracker/internal/config"
)

// LogEntry 日志条目
type LogEntry struct {
	Time    string `json:"time"`
	Content string `json:"content"`
	Exp     int    `json:"exp"`
}

// Storage 数据存储
type Storage struct {
	dataDir string
}

// New 创建存储实例
func New() *Storage {
	configDir := config.GetConfigDir()
	return &Storage{
		dataDir: configDir,
	}
}

// getLogPath 获取日志文件路径
func (s *Storage) getLogPath() string {
	return filepath.Join(s.dataDir, "logs.json")
}

// AddLog 添加日志
func (s *Storage) AddLog(content string, exp int) error {
	logPath := s.getLogPath()
	
	// 读取现有日志
	var logs []LogEntry
	if data, err := os.ReadFile(logPath); err == nil {
		json.Unmarshal(data, &logs)
	}
	
	// 添加新日志
	entry := LogEntry{
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Content: content,
		Exp:     exp,
	}
	logs = append(logs, entry)
	
	// 保存日志
	data, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化日志失败: %w", err)
	}
	
	if err := os.WriteFile(logPath, data, 0644); err != nil {
		return fmt.Errorf("写入日志失败: %w", err)
	}
	
	// 更新配置
	cfg, _ := config.Load()
	if cfg != nil {
		cfg.LogCount++
		cfg.AddExp(exp)
		cfg.Save()
	}
	
	return nil
}

// GetLogs 获取所有日志
func (s *Storage) GetLogs() ([]LogEntry, error) {
	logPath := s.getLogPath()
	
	data, err := os.ReadFile(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []LogEntry{}, nil
		}
		return nil, fmt.Errorf("读取日志失败: %w", err)
	}
	
	var logs []LogEntry
	if err := json.Unmarshal(data, &logs); err != nil {
		return nil, fmt.Errorf("解析日志失败: %w", err)
	}
	
	return logs, nil
}

// GetTodayLogs 获取今日日志
func (s *Storage) GetTodayLogs() ([]LogEntry, error) {
	logs, err := s.GetLogs()
	if err != nil {
		return nil, err
	}
	
	today := time.Now().Format("2006-01-02")
	var todayLogs []LogEntry
	
	for _, log := range logs {
		if len(log.Time) >= 10 && log.Time[:10] == today {
			todayLogs = append(todayLogs, log)
		}
	}
	
	return todayLogs, nil
}
