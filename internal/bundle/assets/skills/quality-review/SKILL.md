---
name: quality-review
description: 检查 SoloWeave 项目的代码改动、工程配置、测试结果，或判断项目是否可以交接与交付时使用。
slug: shiyuandi-soloweave-quality-review
version: 0.1.2
displayName: SoloWeave 质量检查
summary: 核对代码、配置与真实测试结果，并明确报告失败或未执行的检查。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 质量检查

检查实际改动、已确认的 ADR 或工程配置及相关测试。CLI 可用时运行 `soloweave check`；遇到配置或安装问题时使用 `soloweave doctor`。没有 CLI 时，人工核对项目目标、技术决策、交接记录与 Git/源文件是否一致，并运行项目自身适用的测试；将 CLI 检查标为“工具不可用”，不要伪称已运行。根据改动内容，检查是否出现架构偏移、交接信息缺失或新增的重复代码。

分别列出通过、失败、未执行和工具不可用的检查，并给出实际命令或检查依据。不要把未运行的测试说成已通过，也不要把人工审阅称作 CLI 校验。报告具体问题、受影响文件和可执行的下一步。切换 Agent 前，将核实过的结果写入交接记录。
