# 当前状态

核对日期：2026-10-11。

## 当前任务

按开发者要求将仓库文档统一为新的 Skills 项目：删除过时产品设计、迁移说明、配置归档和架构记录；中英文公开说明只介绍当前产品。已写入当前开发计划及 ADR-0004，六个 Skill 内容与版本不变，真实回测证据保留。

开发者已明确授权中文提交并推送本轮文档整理。实际提交与远端状态以 Git 核对为准。本轮 49 份 Markdown 的非空与相对文件链接检查、六项 Skill 资源检查、Git 空白检查通过。前一轮的线上 CI 不代表本轮未提交文档已经通过远端 CI。

## 已有交付基础

- 六项版本：soloweave / feature-workflow 0.2.3，project-setup 0.3.2，project-continuity 0.2.2，quality-review 0.2.1，systematic-debugging 0.1.0。
- 15 项维护回归、确定性打包及资源字节核对已通过；12 个真实 Codex 会话通过六项独立最小场景、空目录路由、受阻双会话及获准修复。
- skills@1.7.1 完成独立安装与公开 GitHub 安装核对；公开安装选择 Codex/Claude Code/Cursor，两套目录共 14 个文件与源码一致。
- 交付记录提交 07e6288 已推送，远端 CI 通过。SkillHub 六项审核已提交；当时交接已发布，其余五项安全审核中。见 docs/release-status.md，后续状态以平台为准。
- 跨设备 Claude Code CLI 原始证据实际模型为 glm-5.3，且使用较早 Skill 修订；不可视为 Anthropic 模型或当前所有版本通过。

## 后续

核对 SkillHub 审核与公开下载；在可用环境单独验证 Cursor、Anthropic 模型、更多项目类型及已有安装冲突。辅助 quick_validate.py 因缺 PyYAML 未成功。不重新全局安装开发者已删除的 Skills。
