# 项目交接

## 当前任务

完成 V0.1 双语公开文档、仓库整理和首次本地提交，并按用户指定修正提交作者。

## 当前状态

中文是公开文档主版本；`docs/en/` 提供 README、使用说明、目录说明及贡献类文档的十组英文对应版本。v0.2 设计基准与历史开发计划按用户选择保留中文原稿。`.gitignore` 排除本地工具和预览构建，`.gitattributes` 统一文本换行。本仓库 Git 作者使用用户明确提供的 `ShiYuandi <shiyuandi@foxmail.com>`。

## 最近改动

- 增加中英文双向入口并核对相对链接。
- 清理冗余 Go 下载包，保留被忽略的本地工具链和预览构建。
- 完善仓库目录说明、贡献文档及项目状态记录。

## 实际验证

- 十组公开文档的英文对应文件与 Markdown 相对链接检查：通过。
- `go test ./...`：通过。
- `go vet ./...`：通过。
- `go build ./...`：通过。
- `git diff --cached --check`：通过。

## 已知限制

- Claude Code、Cursor 的真实客户端交接及 Linux、macOS 预览程序的实际运行尚未验收。
- 自动审批拒绝递归删除旧 `.tools/smoke-project/`；它已被 Git 忽略，不进入提交。
- 本检查点保存于首次提交前。提交改变 Git HEAD 后，`context check` 可能报告过期；请以 `git log`、`git status` 和实际文件为准。

## 下一步

按 `docs/agent-acceptance.md` 完成真实客户端验收；由仓库所有者决定何时推送、运行远程 CI 和发布 Release。

## Git 状态

检查点元数据保存于首次提交前，已因提交变化过期。当前提交及作者请以 `git log -1` 核对，工作区以 `git status` 为准。
