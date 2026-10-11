# 从旧 Go CLI 迁移

[English](en/cli-migration.md)

当前 SoloWeave 已停止维护并移除 Go CLI。直接安装六个 Skills 即可使用，无需下载可执行程序或安装替代 CLI。

历史 [v0.1.1-preview Release](https://github.com/ShiYuandi/SoloWeave/releases/tag/v0.1.1-preview) 保留原下载与对应标签源码；其中五个内置 Skill 是旧版快照。本地迁移没有修改线上 Release。

## 已有项目的迁移步骤

1. 先备份目标项目，核对 README、旧 YAML 配置、批准记录、ADR、交接文档与实际实现。
2. 将仍有效的目标、入口和已确认决策保留在已有 Markdown 或 PROJECT.md 中。保留确认历史；新架构需要重新确认，不能直接沿用旧批准状态。
3. 更新 STATUS.md 和 HANDOFF.md，写清当前工作、真实验证、阻塞和下一步。
4. 安装前检查项目及全局的同名 Skills，保留用户修改并审阅合并。使用标准 Skills 安装器或让 AI 安装完整目录；旧 CLI 安装器不能升级这套新 Skills。
5. 确认项目工作流已不依赖旧机器文件后，再将其归档。迁移不会自动删除其他项目的配置、安装目录或用户文件。

## 原命令现在如何处理

| 原 CLI 功能 | 当前做法 |
| --- | --- |
| init / approve | 用项目设置 Skill 讨论选型，由开发者确认后记录到 PROJECT 和 ADR。 |
| install | 使用标准 Skills 安装器或让 AI 安装，先检查同名文件与用户修改。 |
| context show / checkpoint / resume | 读取、核对并维护 Markdown 目标、状态、变更和交接。 |
| context check / check / doctor | 根据项目实际文件与已有检查命令进行核对，并报告真实结果。 |
| version | 查看各 SKILL.md 的独立版本。 |

Skills 不提供旧 CLI 的机器批准、安装清单或检查点过期检测保证；Agent 不能把人工核对描述为机器校验通过。

本仓库旧机器记录保存在[历史目录](archive/go-cli/README.md)。旧[配置说明](archive/go-cli/configuration.md)和[连续性说明](archive/go-cli/project-continuity.md)仅供查阅历史。
