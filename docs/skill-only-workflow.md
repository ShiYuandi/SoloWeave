# 只安装 Skills 使用 SoloWeave

[English](en/skill-only-workflow.md)

这一方式只需要五个 SoloWeave Skills，不需要下载 `soloweave.exe`。AI 可以遵守工程规范，并用普通项目文件记录已确认的决策和交接信息。它不会获得 CLI 的自动配置校验、安装记录或检查点过期检测。

## 开始使用

按 [README 的提示词](../README.md#让-ai-安装-skills)安装全部五个 Skills，并让 AI 报告每个 `SKILL.md` 的安装路径与版本。打开目标项目后，对 AI 说：“请按 SoloWeave 的规范继续这个项目。先查看现有代码和项目记录；没有 CLI 时使用纯 Skills 流程，明确说明未执行的 CLI 检查。”

## 规划与确认

AI 先看项目已有的 README、依赖文件、代码和 ADR，不覆盖已有约定。需要新建记录时，在 `.soloweave/context/PROJECT.md` 写目标、用户、主要模块、技术选择、运行方式和待确认的问题。技术选择在你明确确认前标记为“待开发者确认”。

你确认后，AI 才在 `.soloweave/decisions/ADR-NNNN.md` 写日期、备选方案及取舍、最终决定、理由和确认事实，并从 `PROJECT.md` 链接该 ADR。变更重要选型时再取得确认并新增 ADR。人工 ADR 是可供下一位 AI 阅读的决策记录，**不等于 CLI 的批准状态**；不要手工创建 `approval.json`，也不要把 `project.yaml` 改成 `approved` 来冒充 CLI 审批。

## 开发、检查与交接

开发前让 AI 检查现有实现、可复用组件和测试。完成改动后运行项目自身适用的测试；报告每项检查的实际命令和结果。没运行的写“未运行”，没有 CLI 时把 `soloweave check` 标为“工具不可用”，不要写成通过。

在需要交接时，让 AI 更新 `.soloweave/context/` 下的文件；已有同类文档时先使用已有文件，不重复建档：

| 文件 | 至少记录 |
| --- | --- |
| `PROJECT.md` | 项目目标、已确认决策及 ADR、主要入口与运行方式 |
| `STATUS.md` | 已完成、进行中、受阻和下一步 |
| `CHANGES.md` | 日期、重要改动及涉及的文件 |
| `HANDOFF.md` | 当前任务、最近改动、实际验证与未运行项、问题、下一步、核对时间和 Git 状态 |

Git 可用时，核对并记录当前分支、HEAD 和工作区变化；不可用时明确写“Git 不可用”。不要把密码、令牌、私钥写入交接文件。新 AI 接手时先读这些记录，再核对源文件与 Git；两者不一致时以实际文件为准，指出记录可能过期。纯 Skills 模式不会生成 `checkpoint.json`，也没有 CLI 的自动过期判断。

## 以后想使用 CLI

Windows x64 可从 [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) 下载可运行 ZIP，无需自行构建。先核对已有文档，再按 [README 的 CLI 路线](../README.md#快速开始)初始化、明确批准并预览安装。CLI 不会自动把人工 ADR 转成它的 `approval.json`；发现同名文件时先处理冲突，不要覆盖已有工作。
