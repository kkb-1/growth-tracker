package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfigDir(t *testing.T) {
	dir := GetConfigDir()
	if dir == "" {
		t.Error("GetConfigDir 返回空字符串")
	}
	
	// 应该包含 .growth-tracker
	if !contains(dir, ".growth-tracker") {
		t.Errorf("配置目录应该包含 .growth-tracker: %s", dir)
	}
}

func TestInitDefault(t *testing.T) {
	// 使用临时目录
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	
	cfg, err := InitDefault()
	if err != nil {
		t.Fatalf("InitDefault 失败: %v", err)
	}
	
	// 验证默认值
	if cfg.Username != "修行者" {
		t.Errorf("Username 期望 '修行者', 得到 %s", cfg.Username)
	}
	
	if len(cfg.Skills) != 4 {
		t.Errorf("期望 4 个技能, 得到 %d", len(cfg.Skills))
	}
	
	// 验证文件是否创建
	configPath := filepath.Join(tmpDir, ".growth-tracker", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error("配置文件未创建")
	}
}

func TestConfigSaveAndLoad(t *testing.T) {
	// 使用临时目录
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	
	// 初始化
	cfg, err := InitDefault()
	if err != nil {
		t.Fatalf("InitDefault 失败: %v", err)
	}
	
	// 修改配置
	cfg.Username = "TestUser"
	cfg.TotalExp = 50
	
	// 保存
	if err := cfg.Save(); err != nil {
		t.Fatalf("Save 失败: %v", err)
	}
	
	// 重新加载
	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load 失败: %v", err)
	}
	
	// 验证
	if loaded.Username != "TestUser" {
		t.Errorf("Username 期望 'TestUser', 得到 %s", loaded.Username)
	}
	
	if loaded.TotalExp != 50 {
		t.Errorf("TotalExp 期望 50, 得到 %d", loaded.TotalExp)
	}
}

func TestAddExp(t *testing.T) {
	cfg := &Config{
		Skills: []Skill{
			{Name: "Go语言", Level: 1, MaxLevel: 10},
		},
		TotalExp: 0,
	}
	
	// 添加经验
	cfg.AddExp(50)
	if cfg.TotalExp != 50 {
		t.Errorf("TotalExp 期望 50, 得到 %d", cfg.TotalExp)
	}
	
	// 添加更多经验触发升级
	cfg.AddExp(60) // 总计 110，应该升到 Lv.2
	if cfg.TotalExp != 110 {
		t.Errorf("TotalExp 期望 110, 得到 %d", cfg.TotalExp)
	}
	
	if cfg.Skills[0].Level != 2 {
		t.Errorf("Go语言等级期望 2, 得到 %d", cfg.Skills[0].Level)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
