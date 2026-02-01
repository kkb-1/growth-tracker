package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "growth-tracker",
	Short: "追踪程序员的成长轨迹",
	Long: `🌱 Growth Tracker - 让每一步进步都被记录
n
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
	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(goalCmd)
	rootCmd.AddCommand(skillCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化配置",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 初始化 Growth Tracker...")
		
		// 创建配置目录
		configDir, _ := os.UserHomeDir()
		configPath := configDir + "/.growth-tracker"
		
		if err := os.MkdirAll(configPath, 0755); err != nil {
			fmt.Printf("❌ 创建目录失败: %v\n", err)
			return
		}
		
		// 创建初始配置文件
		configFile := configPath + "/config.yaml"
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			content := `# Growth Tracker 配置
username: "Your Name"
start_date: "2025-01-01"

# 阶段目标
goals:
  - name: "成为最强的程序员"
    language: "Go"
    status: "进行中"

# 技能树
skills:
  - name: "Go语言"
    level: 1
    max_level: 10
`
			if err := os.WriteFile(configFile, []byte(content), 0644); err != nil {
				fmt.Printf("❌ 创建配置失败: %v\n", err)
				return
			}
		}
		
		fmt.Printf("✅ 初始化完成！配置文件: %s\n", configFile)
		fmt.Println("📖 开始记录你的成长之旅吧！")
	},
}

var logCmd = &cobra.Command{
	Use:   "log [内容]",
	Short: "记录今日学习",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		content := ""
		for _, arg := range args {
			content += arg + " "
		}
		
		fmt.Printf("📝 已记录: %s\n", content)
		fmt.Println("✨ 每一天的进步都值得被记住！")
		
		// TODO: 保存到日志文件
	},
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "查看成长统计",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📊 Growth Tracker 统计")
		fmt.Println("====================")
		fmt.Println()
		fmt.Println("🎯 当前阶段目标: 成为最强的程序员")
		fmt.Println("🔧 主修语言: Go")
		fmt.Println()
		fmt.Println("📈 今日数据:")
		fmt.Println("  - 学习记录: 3 条")
		fmt.Println("  - 代码行数: 150 行")
		fmt.Println("  - 技能提升: +2 exp")
		fmt.Println()
		fmt.Println("🌳 技能树:")
		fmt.Println("  Go语言      [█░░░░░░░░░] 1/10")
		fmt.Println("  系统设计    [░░░░░░░░░░] 0/10")
		fmt.Println("  开源贡献    [░░░░░░░░░░] 0/10")
	},
}

var goalCmd = &cobra.Command{
	Use:   "goal",
	Short: "管理阶段目标",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🎯 阶段目标管理")
		fmt.Println("================")
		fmt.Println()
		fmt.Println("当前目标: 成为最强的程序员")
		fmt.Println("状态: 🌱 进行中")
		fmt.Println()
		fmt.Println("子任务:")
		fmt.Println("  [x] 搭建Go开发环境")
		fmt.Println("  [ ] 完成第一个Go项目")
		fmt.Println("  [ ] 发布到GitHub")
		fmt.Println("  [ ] 获得第一个Star")
	},
}

var skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "查看技能树",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🌳 技能树")
		fmt.Println("==========")
		fmt.Println()
		fmt.Println("核心技能:")
		fmt.Println("  🎯 Go语言精通    Lv.1 → Lv.10")
		fmt.Println("  🏗️  系统设计      Lv.0 → Lv.10")
		fmt.Println("  📦 开源维护      Lv.0 → Lv.10")
		fmt.Println("  🤖 Prompt工程    Lv.0 → Lv.10")
		fmt.Println()
		fmt.Println("经验值: 15 / 100 (升级还需 85 exp)")
	},
}
