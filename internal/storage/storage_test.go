package storage

import (
	"os"
	"testing"

	"github.com/codex/growth-tracker/internal/config"
)

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Error("New() 返回 nil")
	}
	if s.dataDir == "" {
		t.Error("dataDir 为空")
	}
}

func TestAddLog(t *testing.T) {
	// 使用临时目录
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	
	// 先初始化配置
	_, err := config.InitDefault()
	if err != nil {
		t.Fatalf("InitDefault 失败: %v", err)
	}
	
	s := New()
	
	// 添加日志
	err = s.AddLog("测试日志内容", 5)
	if err != nil {
		t.Fatalf("AddLog 失败: %v", err)
	}
	
	// 验证日志是否保存
	logs, err := s.GetLogs()
	if err != nil {
		t.Fatalf("GetLogs 失败: %v", err)
	}
	
	if len(logs) != 1 {
		t.Errorf("期望 1 条日志, 得到 %d", len(logs))
	}
	
	if logs[0].Content != "测试日志内容" {
		t.Errorf("日志内容不匹配: %s", logs[0].Content)
	}
	
	if logs[0].Exp != 5 {
		t.Errorf("日志经验值期望 5, 得到 %d", logs[0].Exp)
	}
}

func TestGetLogsEmpty(t *testing.T) {
	// 使用临时目录
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	
	// 初始化配置
	_, _ = config.InitDefault()
	
	s := New()
	
	// 空日志
	logs, err := s.GetLogs()
	if err != nil {
		t.Fatalf("GetLogs 失败: %v", err)
	}
	
	if len(logs) != 0 {
		t.Errorf("期望空日志列表, 得到 %d 条", len(logs))
	}
}

func TestGetTodayLogs(t *testing.T) {
	// 使用临时目录
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	
	// 初始化配置
	_, _ = config.InitDefault()
	
	s := New()
	
	// 添加今日日志
	s.AddLog("今日学习", 5)
	
	// 获取今日日志
	logs, err := s.GetTodayLogs()
	if err != nil {
		t.Fatalf("GetTodayLogs 失败: %v", err)
	}
	
	if len(logs) != 1 {
		t.Errorf("期望 1 条今日日志, 得到 %d", len(logs))
	}
}
