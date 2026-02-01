package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	Username   string     `yaml:"username"`
	StartDate  string     `yaml:"start_date"`
	Goals      []Goal     `yaml:"goals"`
	Skills     []Skill    `yaml:"skills"`
	LogCount   int        `yaml:"log_count"`
	CodeLines  int        `yaml:"code_lines"`
	TotalExp   int        `yaml:"total_exp"`
}

// Goal 阶段目标
type Goal struct {
	Name   string `yaml:"name"`
	Lang   string `yaml:"language"`
	Status string `yaml:"status"`
}

// Skill 技能
type Skill struct {
	Name     string `yaml:"name"`
	Level    int    `yaml:"level"`
	MaxLevel int    `yaml:"max_level"`
}

// GetConfigDir 获取配置目录
func GetConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".growth-tracker")
}

// GetConfigPath 获取配置文件路径
func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.yaml")
}

// Load 加载配置
func Load() (*Config, error) {
	configPath := GetConfigPath()
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}
	
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	
	return &cfg, nil
}

// Save 保存配置
func (c *Config) Save() error {
	configPath := GetConfigPath()
	
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %w", err)
	}
	
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("写入配置失败: %w", err)
	}
	
	return nil
}

// InitDefault 初始化默认配置
func InitDefault() (*Config, error) {
	configDir := GetConfigDir()
	
	// 创建目录
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("创建配置目录失败: %w", err)
	}
	
	cfg := &Config{
		Username:  "修行者",
		StartDate: "2025-01-01",
		Goals: []Goal{
			{
				Name:   "成为最强的程序员",
				Lang:   "Go",
				Status: "进行中",
			},
		},
		Skills: []Skill{
			{Name: "Go语言", Level: 1, MaxLevel: 10},
			{Name: "系统设计", Level: 0, MaxLevel: 10},
			{Name: "开源维护", Level: 0, MaxLevel: 10},
			{Name: "Prompt工程", Level: 0, MaxLevel: 10},
		},
		LogCount:  0,
		CodeLines: 0,
		TotalExp:  0,
	}
	
	if err := cfg.Save(); err != nil {
		return nil, err
	}
	
	return cfg, nil
}

// AddExp 增加经验值
func (c *Config) AddExp(exp int) {
	c.TotalExp += exp
	
	// 检查技能升级 (每100exp升1级，初始Lv.1)
	for i := range c.Skills {
		if c.Skills[i].Name == "Go语言" {
			newLevel := 1 + c.TotalExp/100
			if newLevel > c.Skills[i].Level && newLevel <= c.Skills[i].MaxLevel {
				c.Skills[i].Level = newLevel
			}
		}
	}
}
