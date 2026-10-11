# 交付状态与历史快照

## 当前五项版本：GitHub 更新

2026-10-11 开发者确认移除独立故障排查。当前源码为五项，入口和功能开发为 0.2.4；本轮验证见[验收记录](agent-acceptance.md)。本地验收完成后，开发者已明确授权中文提交并推送至 GitHub；实际提交、远端与 CI 结果按 Git/GitHub 核对。本轮未更新 SkillHub，也未卸载用户已有安装。旧 SkillHub 条目需在取得平台更新授权后核对处理。

## 历史六项交付快照

以下内容记录此前六项交付，不代表当前五项已经发布；版本和测试结果按当时事实保留。


日期：2026-10-11。开发者明确授权的提交、推送、GitHub 验证和六项 SkillHub 审核提交已完成。

### 当时本地验收

- 六个独立 Skill、维护脚本和双语使用文档完成。
- 15 项静态回归检查通过；六包内容、SHA256SUMS 和重复打包核对通过。
- 当时候选版本的六项独立最小行为、全套空目录路由、受阻接手及获准修复通过；12 个真实 Codex 会话的边界见[结果](codex-behavior-results.md)。

### 当时线上步骤

| 项目 | 当前状态 |
| --- | --- |
| 中文提交与推送 | 已推送至 origin/master，交付记录提交为 `07e6288`。 |
| GitHub CI | [135e13c 检查通过](https://github.com/ShiYuandi/SoloWeave/actions/runs/38146932426)。首次检查的 15 项全部通过，但负例进程残留退出码导致失败；已明确返回成功退出码并经远端复验。 |
| 公开 GitHub 安装 | skills@1.7.1 从 ShiYuandi/SoloWeave 成功安装六项到全新临时项目，选择 Codex、Claude Code、Cursor；两套目录共 14 个文件与源码字节一致。未改动用户全局安装。 |
| SkillHub | 五项旧 Skill 已上传对应源码包，故障排查通过 GitHub 导入新增；六项均已提交审核。项目设置包含 selection-guide.md。 |

#### 当时 SkillHub 提交后状态

2026-10-11 管理页核对结果；这是当时快照，后续以平台页面为准。

| Skill | 版本 | 状态 |
| --- | --- | --- |
| soloweave | 0.2.3 | 安全审核中 |
| project-setup | 0.3.2 | 安全审核中 |
| feature-workflow | 0.2.3 | 安全审核中 |
| project-continuity | 0.2.2 | 已发布 |
| quality-review | 0.2.1 | 安全审核中 |
| systematic-debugging | 0.1.0 | 安全审核中 |

管理页可能在审核中继续显示旧名称；入口和设置已提交源码中的中文名称与概述。更新表单使用 ZIP，新建调试项使用 GitHub 导入。截图保存在本地忽略目录 `dist/acceptance/skillhub-submitted-2026-10-11.jpg`。

提交审核不等于审核通过；五项待审版本的公开下载尚未验收。当前版本、链接和发布步骤见 [Skills](skills.md)与[发布指南](releasing.md)。未创建新 GitHub Release。
