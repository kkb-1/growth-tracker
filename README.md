# Growth Tracker 🌱

[![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

> *追踪程序员的成长轨迹，让每一步进步都被记录。*

## ✨ 功能特性

- 📅 **每日记录** - 快速记录今日所学，自动保存
- 🎯 **阶段目标** - 设定并追踪长期修行目标
- 🌳 **技能树** - 可视化技能掌握程度与升级进度
- 📊 **数据统计** - 生成成长报告与经验值统计
- 📝 **历史查看** - 查看今日/历史学习记录

## 📦 安装

### 方式一：go install
```bash
go install github.com/codex/growth-tracker@latest
```

### 方式二：源码编译
```bash
git clone https://github.com/codex/growth-tracker.git
cd growth-tracker
go build -o growth-tracker .
```

## 🚀 快速开始

```bash
# 1. 初始化配置
growth-tracker init

# 2. 记录今日学习
growth-tracker log "学习了Go的接口和结构体"
growth-tracker log "完成了数据持久化功能"

# 3. 查看今日记录
growth-tracker list

# 4. 查看成长统计
growth-tracker stats

# 5. 查看技能树
growth-tracker skill
```

## 📖 命令详解

| 命令 | 说明 | 示例 |
|------|------|------|
| `init` | 初始化配置 | `growth-tracker init` |
| `log` | 记录学习 | `growth-tracker log "今天学了xxx"` |
| `list` | 查看今日记录 | `growth-tracker list` |
| `stats` | 成长统计 | `growth-tracker stats` |
| `goal` | 阶段目标 | `growth-tracker goal` |
| `skill` | 技能树 | `growth-tracker skill` |

## 🧪 开发

```bash
# 运行测试
go test ./... -v

# 编译
go build -o growth-tracker .

# 安装到 $GOPATH/bin
go install
```

## 📁 项目结构

```
growth-tracker/
├── cmd/
│   └── root.go          # CLI命令实现
├── internal/
│   ├── config/          # 配置管理
│   │   ├── config.go
│   │   └── config_test.go
│   └── storage/         # 数据存储
│       ├── storage.go
│       └── storage_test.go
├── main.go              # 入口
├── go.mod
├── go.sum
└── README.md
```

## 🎯 Roadmap

- [x] 基础CLI框架
- [x] 数据持久化
- [x] 单元测试
- [ ] 配置文件编辑
- [ ] 数据导出(JSON/CSV)
- [ ] GitHub贡献集成
- [ ] 多用户支持

## 📄 License

MIT License

---

> *Created with 🦾 by Codex - 最强程序员修行之旅 Day 1*

[⬆ Back to Top](#growth-tracker-)
