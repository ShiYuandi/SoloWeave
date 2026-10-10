---
name: quality-review
description: 检查 SoloWeave 项目的代码改动、工程配置、测试结果，或判断项目是否可以交接与交付时使用。
slug: shiyuandi-soloweave-quality-review
version: 0.1.1
displayName: SoloWeave 质量检查
summary: 核对代码、配置与真实测试结果，并明确报告失败或未执行的检查。
license: MIT
homepage: https://github.com/ShiYuandi/SoloWeave
---

# 质量检查

检查实际改动和相关工程配置。运行适用的测试与 `soloweave check`；遇到配置或安装问题时，使用 `soloweave doctor`。根据改动内容，检查是否出现架构偏移、交接信息缺失或新增的重复代码。

分别列出通过、失败、未执行和工具不可用的检查。不要把未运行的测试说成已通过。报告具体问题、受影响文件和可执行的下一步。切换 Agent 前，将核实过的结果写入交接记录。
