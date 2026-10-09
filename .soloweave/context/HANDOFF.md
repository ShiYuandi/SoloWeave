# 项目交接

## 当前任务

简化双语文档的语言切换标注，并记录中文 commit 信息约定。

## 当前状态

中英文公开文档保留双向链接，不再显示语言优先级标注。`AGENTS.md` 和中英文贡献说明记录了后续 commit 信息使用中文的约定。本仓库 Git 作者为 `ShiYuandi <shiyuandi@foxmail.com>`。

## 最近改动

- 清理中英文文档中的语言优先级标注，保留双向入口。
- 更新仓库指引和中英文贡献说明，明确中文 commit 信息约定。

## 实际验证

- 十组公开文档的英文对应文件与 Markdown 相对链接检查：通过。
- `go test ./...`：通过。
- `go vet ./...`：通过。
- `go build ./...`：通过。
- `git diff --cached --check`：通过。
- 本轮 Markdown 相对链接与 `git diff --check`：通过；本轮未修改 Go 代码。

## 已知限制

- Claude Code、Cursor 的真实客户端交接及 Linux、macOS 预览程序的实际运行尚未验收。
- 自动审批拒绝递归删除旧 `.tools/smoke-project/`；它已被 Git 忽略，不进入提交。
- 本检查点保存于首次提交前。提交改变 Git HEAD 后，`context check` 可能报告过期；请以 `git log`、`git status` 和实际文件为准。

## 下一步

按 `docs/agent-acceptance.md` 完成真实客户端验收；由仓库所有者决定何时推送、运行远程 CI 和发布 Release。

## Git 状态

检查点元数据保存于首次提交前，已因提交变化过期。当前提交及作者请以 `git log -1` 核对，工作区以 `git status` 为准。
