#!/bin/bash
# GitHub自动发布脚本

echo "🚀 Growth Tracker GitHub自动发布脚本"
echo "======================================"
echo

# 检查GitHub用户名
GITHUB_USER="kkb"
REPO_NAME="growth-tracker"

echo "📦 项目信息:"
echo "  用户: $GITHUB_USER"
echo "  仓库: $REPO_NAME"
echo

# 检查是否有GitHub Token
if [ -z "$GITHUB_TOKEN" ]; then
    echo "⚠️  需要GitHub Token来创建仓库"
    echo
    echo "获取Token步骤:"
    echo "1. 访问: https://github.com/settings/tokens/new"
    echo "2. 选择 scopes: repo, workflow"
    echo "3. 生成Token并复制"
    echo
    echo "然后运行: export GITHUB_TOKEN=你的token"
    echo "再运行此脚本"
    exit 1
fi

echo "✅ GitHub Token已配置"
echo

# 创建仓库
echo "📁 创建GitHub仓库..."
curl -s -X POST \
  -H "Authorization: token $GITHUB_TOKEN" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/user/repos \
  -d "{\"name\":\"$REPO_NAME\",\"description\":\"🌱 追踪程序员的成长轨迹 - CLI工具\",\"private\":false}" > /tmp/repo_create.json

if grep -q "\"id\":" /tmp/repo_create.json; then
    echo "✅ 仓库创建成功!"
    REPO_URL="git@github.com:$GITHUB_USER/$REPO_NAME.git"
    echo "   URL: https://github.com/$GITHUB_USER/$REPO_NAME"
elif grep -q "name already exists" /tmp/repo_create.json; then
    echo "⚠️  仓库已存在，使用现有仓库"
    REPO_URL="git@github.com:$GITHUB_USER/$REPO_NAME.git"
else
    echo "❌ 创建失败:"
    cat /tmp/repo_create.json
    exit 1
fi

echo

# 配置git remote
cd /Users/bkwang/.openclaw/workspace/projects/growth-tracker

echo "🔗 配置Git远程仓库..."
if git remote | grep -q origin; then
    git remote remove origin
fi
git remote add origin $REPO_URL

echo "📤 推送到GitHub..."
git push -u origin main

if [ $? -eq 0 ]; then
    echo
    echo "🎉 发布成功!"
    echo "=============="
    echo "仓库地址: https://github.com/$GITHUB_USER/$REPO_NAME"
    echo
    echo "安装命令:"
    echo "  go install github.com/$GITHUB_USER/$REPO_NAME@latest"
else
    echo "❌ 推送失败"
    exit 1
fi
