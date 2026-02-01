package cmd

import (
	"fmt"
	"strings"

	"github.com/codex/growth-tracker/internal/config"
	"github.com/codex/growth-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "growth-tracker",
	Short: "追踪程序员的成长轨迹",
	Long: `🌱 Growth Tracker - 让每一步进步都被记录

一个帮助开发者记录每日学习、设定阶段目标、
追踪技能树进化的命令行工具。`,
	Version: "0.1.0",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(goalCmd)
	rootCmd.AddCommand(skillCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化配置",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🚀 初始化 Growth Tracker...")
		
		cfg, err := config.InitDefault()
		if err != nil {
			return fmt.Errorf("❌ 初始化失败: %w", err)
		}
		
		fmt.Printf("✅ 初始化完成！\n")
		fmt.Printf("📁 配置目录: %s\n", config.GetConfigDir())
		fmt.Printf("👤 修行者: %s\n", cfg.Username)
		fmt.Println("📖 开始记录你的成长之旅吧！")
		fmt.Println()
		fmt.Println("💡 快速开始:")
		fmt.Println("  growth-tracker log \"今天学习了...\"")
		fmt.Println("  growth-tracker stats")
		
		return nil
	},
}

var logCmd = &cobra.Command{
	Use:   "log [内容]",
	Short: "记录今日学习",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content := strings.Join(args, " ")
		
		// 保存到存储
		store := storage.New()
		if err := store.AddLog(content, 5); err != nil {
			return fmt.Errorf("❌ 保存失败: %w", err)
		}
		
		fmt.Printf("📝 已记录: %s\n", content)
		fmt.Println("✨ 每一天的进步都值得被记住！")
		
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "查看今日记录",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.New()
		logs, err := store.GetTodayLogs()
		if err != nil {
			return fmt.Errorf("❌ 读取失败: %w", err)
		}
		
		if len(logs) == 0 {
			fmt.Println("📭 今日暂无记录")
			fmt.Println("💡 使用: growth-tracker log \"今天学习了...\"")
			return nil
		}
		
		fmt.Println("📋 今日学习记录")
		fmt.Println("================")
		fmt.Println()
		
		for i, log := range logs {
			fmt.Printf("%d. [%s] %s (+5 exp)\n", i+1, log.Time[11:16], log.Content)
		}
		
		fmt.Printf("\n📊 今日共 %d 条记录\n", len(logs))
		
		return nil
	},
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "查看成长统计",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("❌ 请先运行: growth-tracker init")
		}
		
		store := storage.New()
		todayLogs, _ := store.GetTodayLogs()
		totalLogs, _ := store.GetLogs()
		
		fmt.Println("📊 Growth Tracker 统计")
		fmt.Println("====================")
		fmt.Println()
		
		if len(cfg.Goals) > 0 {
			fmt.Printf("🎯 当前阶段目标: %s\n", cfg.Goals[0].Name)
			fmt.Printf("🔧 主修语言: %s\n", cfg.Goals[0].Lang)
			fmt.Printf("📌 状态: %s\n", cfg.Goals[0].Status)
		}
		
		fmt.Println()
		fmt.Println("📈 数据统计:")
		fmt.Printf("  - 总记录数: %d 条\n", len(totalLogs))
		fmt.Printf("  - 今日记录: %d 条\n", len(todayLogs))
		fmt.Printf("  - 累计经验: %d exp\n", cfg.TotalExp)
		
		fmt.Println()
		fmt.Println("🌳 技能树:")
		for _, skill := range cfg.Skills {
			progress := renderProgress(skill.Level, skill.MaxLevel)
			fmt.Printf("  %-10s [%s] %d/%d\n", skill.Name, progress, skill.Level, skill.MaxLevel)
		}
		
		return nil
	},
}

var goalCmd = &cobra.Command{
	Use:   "goal",
	Short: "管理阶段目标",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("❌ 请先运行: growth-tracker init")
		}
		
		fmt.Println("🎯 阶段目标管理")
		fmt.Println("================")
		fmt.Println()
		
		for i, goal := range cfg.Goals {
			fmt.Printf("%d. %s\n", i+1, goal.Name)
			fmt.Printf("   语言: %s | 状态: %s\n", goal.Lang, goal.Status)
			fmt.Println()
		}
		
		fmt.Println("里程碑:")
		milestones := []struct {
			Done bool
			Text string
		}{
			{true, "搭建Go开发环境"},
			{true, "创建growth-tracker项目"},
			{true, "实现数据持久化"},
			{false, "添加单元测试"},
			{false, "发布到GitHub"},
			{false, "获得第一个Star"},
		}
		
		for _, m := range milestones {
			mark := "[ ]"
			if m.Done {
				mark = "[x]"
			}
			fmt.Printf("  %s %s\n", mark, m.Text)
		}
		
		return nil
	},
}

var skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "查看技能树",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("❌ 请先运行: growth-tracker init")
		}
		
		fmt.Println("🌳 技能树")
		fmt.Println("==========")
		fmt.Println()
		fmt.Println("核心技能:")
		
		icons := []string{"🎯", "🏗️", "📦", "🤖"}
		for i, skill := range cfg.Skills {
			icon := icons[i%len(icons)]
			fmt.Printf("  %s %-10s Lv.%d / Lv.%d\n", icon, skill.Name, skill.Level, skill.MaxLevel)
		}
		
		fmt.Println()
		fmt.Printf("💫 总经验值: %d / %d (下一级还需 %d exp)\n", 
			cfg.TotalExp, 
			(cfg.TotalExp/100+1)*100,
			100-(cfg.TotalExp%100))
		
		return nil
	},
}

// renderProgress 渲染进度条
func renderProgress(current, max int) string {
	if max <= 0 {
		return "░░░░░░░░░░"
	}
	
	filled := (current * 10) / max
	if filled > 10 {
		filled = 10
	}
	
	result := ""
	for i := 0; i < 10; i++ {
		if i < filled {
			result += "█"
		} else {
			result += "░"
		}
	}
	
	return result
}
