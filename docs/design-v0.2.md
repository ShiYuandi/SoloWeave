<!-- 本文整理自用户提供的 SoloWeave v0.2 最终整合版。实施计划见 development-plan.md。 -->


# SoloWeave
Build independently. Ship confidently.

面向个人开发者与 1–5 人微型团队的开源 AI 全栈工程工作流与技能包。

## 项目设计文档 v0.2 — 正式开发基准

| 项目属性 | 内容 |
| --- | --- |
| 项目名称 | SoloWeave |
| 项目类型 | 开源 AI 工程技能包与 CLI 工具 |
| 目标用户 | 1–5 人开发者与微型团队 |
| 开发语言 | Go |
| AI 平台 | Codex、Claude Code、Cursor |
| 配置系统 | YAML + JSON Schema |
| 工程能力 | Agent Skills + CLI + Git |
| 当前状态 | 设计完成，准备进入开发 |
| 文档版本 | v0.2 |
| 日期 | 2026-10-09 |

## 1. 项目愿景与定位
### 1.1 项目愿景
SoloWeave 旨在帮助个人开发者和微型团队，更高效、更规范地借助 AI 完成软件项目。

通过统一的工程规范、技术决策、持久化项目上下文和标准化开发流程，降低 AI 辅助开发中的重复沟通、架构漂移、代码重复和上下文丢失问题。

核心理念：开发者负责决策，AI 负责执行，工程规则负责约束，项目上下文负责延续。

### 1.2 主要解决的问题
架构失控
AI 可能未经确认就选择技术栈、改变目录结构或引入复杂架构。

SoloWeave 让开发者决定项目架构，并将决策持久化。

代码重复
AI 可能忽略项目已有的公共组件、接口、工具方法和依赖库。

SoloWeave 要求开发前检查现有实现，合理复用。

开发不连续
更换 AI 账号、切换编程 Agent 或上下文窗口耗尽后，新 AI 不知道项目已经完成了什么。

SoloWeave 将项目状态与交接信息保存到项目仓库中。

开发流程不规范
AI 可能只追求实现当前功能，而忽略需求分析、测试、安全、架构边界和部署要求。

SoloWeave 提供统一的工程工作流和质量检查机制。

平台依赖
不同 AI 编程工具采用不同的规则文件，开发者需要重复维护。

SoloWeave 通过适配层使多个 Agent 使用相同工程约束。

### 1.3 目标用户
SoloWeave 面向：

一人独立开发者。

两至五人的微型开发团队。

使用 AI 从零构建产品的创业者。

需要切换多个 AI 账号或编程 Agent 的开发者。

希望保持工程规范与开发连续性的开源项目维护者。

SoloWeave 应优先满足个人开发体验，同时能够支持小团队通过 Git 进行协作。

### 1.4 非目标
SoloWeave 不计划：

自行训练或托管大语言模型。

替代 Codex、Claude Code、Cursor。

强制开发者采用单一技术栈。

自动保证 AI 生成代码完全正确。

在 V0.1 建立复杂的自主编程 Agent。

在 V0.1 实现图形管理后台。

在 V0.1 依赖 MCP Server 或远程 AI API。

## 2. 三大核心支柱
### 2.1 Engineering Standards — 工程规范
负责让 AI 按照开发者批准的方案持续开发。

主要能力：

技术栈选择

软件架构选择

项目目录规范

API 设计规范

数据库设计规范

代码复用

模块依赖约束

质量检查

架构变更管理

核心原则：

AI 可以推荐，但不能擅自改变已经确认的重要技术决策。

### 2.2 Project Continuity — 项目连续性
负责让不同 AI 能够接续同一个项目的开发工作。

主要能力：

项目介绍持久化

开发进度记录

已完成与未完成任务管理

开发修改记录

任务检查点

AI 交接文档

Git 状态核对

跨账号、跨 Agent 的上下文恢复

核心原则：

聊天上下文可以丢失，但项目的重要决策和开发进度应尽可能保存在项目中。

### 2.3 Development Workflow — 开发工作流
负责规范从项目想法到正式交付的开发过程。

主要能力：

需求分析

MVP 规划

架构设计

前端开发

后端开发

数据库设计

测试与安全检查

部署交付

开发记录及交接

核心原则：

开发任务应当有明确目标、执行过程、验证结果和下一步计划。

## 3. 总体系统架构
SoloWeave 采用平台无关的分层设计。

### 3.1 AI 平台层
支持：

Codex

Claude Code

Cursor

后续通过适配器扩展其他兼容 Agent Skills 或项目规则机制的编程 Agent。

该层负责理解任务、编辑代码和调用工具。

### 3.2 Agent Skills 层
负责 AI 开发流程，包括：

项目初始化

任务分类

架构选型

现有代码分析

代码复用

功能开发

测试与质量审查

项目上下文维护

任务恢复与交接

Skill 只规定 Agent 如何工作，不直接代替项目配置和静态检查工具。

### 3.3 Go 核心引擎
负责确定性工程操作，包括：

CLI 命令

配置读取与校验

架构目录管理

Skill 安装

Agent 平台适配

Git 状态采集

上下文检查点管理

项目扫描

架构规则校验

统一检查报告

Go 核心不依赖某个指定的 AI 模型。

### 3.4 项目工程契约
以 .soloweave/ 目录为统一载体，保存：

技术栈

架构配置

工程规范

开发状态

项目交接信息

重要架构决策

安装元数据

该目录应当能够通过 Git 进行版本管理。

## 4. 正式技术方案

| 领域 | 方案 |
| --- | --- |
| 核心语言 | Go |
| CLI 框架 | Cobra |
| 配置格式 | YAML |
| 配置校验 | JSON Schema + Go |
| 静态文件打包 | Go embed |
| Agent 技能格式 | SKILL.md |
| 开发规则 | AGENTS.md / CLAUDE.md / Cursor Rules |
| 项目上下文 | Markdown + Git |
| 文件处理 | Go 标准库 |
| 外部校验器 | os/exec |
| 自动测试 | Go testing |
| 持续集成 | GitHub Actions |
| 版本管理 | SemVer |
| 发布方式 | 各平台独立可执行程序 |

V0.1 不引入数据库。

工程配置以文件为主要存储形式，便于开发者查看、编辑、审查和版本控制。

## 5. 正式仓库结构
soloweave/
├── cmd/
│   └── soloweave/
│       └── main.go
│
├── internal/
│   ├── cli/
│   ├── config/
│   ├── catalog/
│   ├── architect/
│   ├── project/
│   ├── continuity/
│   ├── installer/
│   ├── adapter/
│   │   ├── codex/
│   │   ├── claude/
│   │   └── cursor/
│   ├── scanner/
│   ├── reuse/
│   ├── validator/
│   ├── report/
│   └── bundle/
│       ├── embed.go
│       └── assets/
│           ├── skills/
│           ├── rules/
│           ├── templates/
│           ├── catalogs/
│           └── schemas/
│
├── docs/
│   ├── vision.md
│   ├── architecture.md
│   ├── configuration.md
│   ├── skills.md
│   ├── project-continuity.md
│   ├── development-workflow.md
│   ├── development-plan.md
│   └── roadmap.md
│
├── testdata/
├── .github/
│   └── workflows/
│
├── AGENTS.md
├── README.md
├── CONTRIBUTING.md
├── SECURITY.md
├── CHANGELOG.md
├── LICENSE
└── go.mod
这是目标目录结构，不要求初始化时创建所有空目录。

continuity/ 负责项目上下文、开发状态和交接信息。

adapter/ 只处理平台差异。

bundle/assets/ 保存内置 Skill、规则和模板的源文件。

scanner/ 与 reuse/ 在后续版本逐步实现。

## 6. 项目配置系统
### 6.1 配置目录
每个采用 SoloWeave 的项目都应保存一套工程配置。

my-project/
├── .soloweave/
│   ├── project.yaml
│   ├── installation.json
│   │
│   ├── context/
│   │   ├── PROJECT.md
│   │   ├── STATUS.md
│   │   ├── HANDOFF.md
│   │   ├── CHANGES.md
│   │   └── tasks/
│   │
│   └── decisions/
│       └── ADR-0001.md
│
├── AGENTS.md
├── CLAUDE.md
├── .agents/
├── .claude/
├── .cursor/
└── src/
### 6.2 配置职责
project.yaml

保存项目的架构、语言、框架、数据库、工程规则和启用的 AI 平台。

installation.json

记录 SoloWeave 的安装版本、启用的 Skill 和平台适配信息。

context/

保存项目介绍、进度、修改记录和交接信息。

decisions/

保存重大架构决策及其原因。

项目技术选型只在 project.yaml 中维护正式值，其他文档引用该配置，避免出现多个互相矛盾的信息来源。

### 6.3 配置示例
schema_version: 1

project:
  name: demo-app
  type: fullstack
  status: approved
  team_size: 1

frontend:
  language: typescript
  framework: nextjs
  package_manager: pnpm

backend:
  language: typescript
  framework: nestjs
  architecture: modular-monolith

database:
  engine: postgresql
  orm: prisma

workspace:
  layout: monorepo
  frontend_path: apps/web
  backend_path: apps/api
  shared_path: packages/shared

reuse:
  inspect_existing: true
  prefer_existing: true
  avoid_premature_abstraction: true

continuity:
  enabled: true
  checkpoint_on_task_completion: true
  checkpoint_on_handoff: true
  verify_git_state: true

quality:
  lint: required
  relevant_tests: required

agents:
  - codex
  - claude
  - cursor
配置支持纯前端、纯后端和全栈项目。

team_size 支持 1–5，但不用于限制实际项目人数。

对关键架构变更，应生成新的 ADR，并由开发者确认。

## 7. 架构选择系统
### 7.1 设计目标
SoloWeave 不默认强制使用某种架构。

开发者可以根据项目情况选择：

前端语言与框架

后端语言与框架

数据库与 ORM

单仓库或多仓库

单体、模块化单体或微服务

分层架构、整洁架构或六边形架构

API 设计方式

认证授权方案

测试和部署方案

领域驱动设计作为可组合的领域建模方法，不需要与所有软件架构互斥。

### 7.2 选型流程
分析项目需求。

确定需要做出的技术决策。

给出合理候选方案。

解释各方案的优缺点。

开发者选择和确认。

检查技术方案兼容性。

生成项目配置和 ADR。

开始工程初始化。

已经批准的架构不应在日常开发中被随意改变。

### 7.3 现有项目
现有项目接入 SoloWeave 时，应当优先识别当前实现。

不得因为接入 SoloWeave 就自动执行大规模架构迁移。

需要先生成候选配置，让开发者确认后再启用相应约束。

## 8. Project Continuity：项目上下文持久化
这是 SoloWeave V0.1 的核心功能。

### 8.1 功能目标
解决以下场景：

更换 AI 账号。

Codex 切换到 Claude Code。

Claude Code 切换到 Cursor。

AI 上下文窗口耗尽。

开发中断后恢复。

将项目交给另一名开发者。

1–5 人团队之间交接任务。

目标是让新 Agent 在不依赖旧聊天记录的情况下，快速理解项目并继续工作。

### 8.2 文档职责

| 文档 | 内容 | 更新时机 |
| --- | --- | --- |
| PROJECT.md | 项目目标、功能范围和系统介绍 | 项目重要信息变化 |
| STATUS.md | 已完成、进行中、待办和阻塞任务 | 任务状态变化 |
| HANDOFF.md | 当前开发位置、下一步、风险和验证结果 | 检查点与交接 |
| CHANGES.md | 已完成的重要代码修改 | 功能或重要修改完成 |
| tasks/ | 各任务的详细目标、状态和进度 | 任务执行过程 |
| ADR | 重大架构决策 | 技术决策发生变化 |

文档应各司其职，不需要每修改一行代码就更新全部文件。

### 8.3 PROJECT.md
主要包含：

项目简介

产品目标

核心功能

已确定的技术方案引用

主要业务模块

开发与运行方式

项目重要入口文件

尽量保持稳定，不保存过多临时开发细节。

### 8.4 STATUS.md
记录项目整体进度。

建议采用：

Completed：已完成

In Progress：进行中

Blocked：阻塞

Planned：待开发

对于已完成事项，应能够定位到对应功能、任务或提交记录。

### 8.5 HANDOFF.md
这是新 AI 恢复开发时首先阅读的文档。

建议包含：

# Project Handoff

## Current Task
当前正在开发的功能

## Current State
已完成与未完成的部分

## Recent Changes
最近修改的模块与文件

## Verification
实际执行的测试与检查结果

## Known Issues
已知问题、阻塞和风险

## Next Steps
下一步需要执行的操作

## Git State
分支、提交标识、未提交修改情况
文档必须保持简洁，避免复制整个历史记录。

新 Agent 应在读取后核对实际源代码和 Git 状态。

### 8.6 AI 开发过程中的记录策略
任务开始

读取配置、项目介绍、当前进度和交接信息。

确认当前工作目标。

开发过程中

在完成重要子任务、产生重大决策或需要暂停时保存检查点。

任务结束

记录：

完成的功能

主要修改文件

重要实现原因

已运行的测试

尚未解决的问题

下一步计划

交接时

生成最新 HANDOFF 信息。

新的 Agent 读取并核对文档后恢复工作。

### 8.7 可靠性要求
不能完全依赖 AI 自己声明开发完成。

Go CLI 应采集可验证的 Git 信息，例如：

当前分支

当前提交

未提交文件变更

受影响文件

最近一次检查点

具体业务进度和修改原因由 Agent 负责描述。

如果源代码和交接文档不一致，应优先核对实际代码和 Git 状态，并提示文档可能过期。

AI 意外中断时，最后一次检查点之后的意图可能丢失，因此不能承诺绝对无损恢复。

### 8.8 多人协作
V0.1 优先保证同一工作区内的串行 Agent 交接。

不同开发者应使用独立分支、独立任务记录，避免多个 Agent 同时覆盖相同状态文件。

后续版本将增强任务级上下文隔离、分支相关检查点和交接记录合并。

### 8.9 安全要求
项目上下文不允许保存：

API Key

密码

Access Token

私钥

真实环境变量密钥

其他无需共享的敏感数据

项目文档同步至 Git 时，必须遵守仓库访问权限和敏感信息保护要求。

## 9. Agent Skills 体系
### 9.1 V0.1 Skills
第一版计划实现五个核心 Skill。

| Skill | 职责 |
| --- | --- |
| soloweave | 主入口、项目规范读取、任务路由 |
| project-setup | 项目规划、架构和技术选型 |
| feature-workflow | 功能开发、代码复用和开发计划 |
| project-continuity | 检查点、开发进度、上下文恢复 |
| quality-review | 测试、校验、质量报告 |

后续增加：

frontend-development

backend-development

database-design

deployment

security-review

V0.1 不需要一次完成所有专业 Skill。

### 9.2 主工作流程
新项目
需求分析 → 技术选型 → 架构确认 → 配置生成 → 项目初始化 → 功能开发 → 测试 → 检查点。

已有项目
读取配置 → 恢复上下文 → 检查代码 → 搜索复用 → 规划修改 → 实施 → 验证 → 更新进度。

切换 Agent
读取交接 → 核对 Git → 定位任务 → 检查已有实现 → 继续开发 → 保存新检查点。

### 9.3 Skill 编写原则
SKILL.md 保持精简。

详细规范放入 references。

通过 assets 提供必要模板。

优先使用项目配置而不是硬编码技术选项。

不要求每次小修改执行完整开发流程。

不重复加载无关参考内容。

不能把未经验证的结果标记为通过。

## 10. Codex / Claude Code / Cursor 适配
### 10.1 核心原则
不同平台使用同一套工程规范和项目配置。

平台适配器只负责发现、安装和加载相应文件。

### 10.2 平台文件
Codex：

AGENTS.md
.agents/skills/
Claude Code：

CLAUDE.md
.claude/skills/
Cursor：

.cursor/rules/
.agents/skills/
实际安装前应依据目标平台版本验证规则和 Skill 加载机制。

### 10.3 安装要求
不覆盖用户现有规则。

支持重复安装。

支持安装预览。

支持安装状态检查。

尽量避免重复保存 Skill 源文件。

使用明确的版本元数据。

确保文件操作不逃逸项目目录。

安装失败时报告状态并避免破坏已有文件。

### 10.4 新 Agent 启动规则
新 Agent 进入项目时，应：

读取项目级规则。

检查 .soloweave/project.yaml。

阅读 context/HANDOFF.md。

阅读相关任务和状态。

核对 Git 与实际代码。

恢复开发。

## 11. CLI 命令规范
正式命令名称：

# SoloWeave

### 11.1 V0.1 命令

| 命令 | 功能 |
| --- | --- |
| soloweave init | 初始化项目配置 |
| soloweave approve | 确认技术与架构决策 |
| soloweave install | 安装 Agent Skills 和规则 |
| soloweave doctor | 检查环境与安装状态 |
| soloweave check | 执行基础工程校验 |
| soloweave version | 输出版本信息 |
| soloweave context show | 查看项目上下文 |
| soloweave context checkpoint | 保存任务检查点 |
| soloweave context resume | 查看恢复信息和 Git 状态 |
| soloweave context check | 检查交接文档完整性和新鲜度 |

这些是计划实现的接口，目前尚非可运行命令。

### 11.2 使用示例
初始化：

soloweave init
安装：

soloweave install --agents codex,claude,cursor
恢复项目：

soloweave context resume
保存检查点：

soloweave context checkpoint
执行检查：

soloweave check
CLI 只负责可验证的数据和文件操作。

检查点的业务摘要应由当前 Agent 或开发者提供，不依赖 CLI 自行调用大模型推断。

## 12. 代码复用系统
### 12.1 核心策略
Reuse Before Reinvent。

创建新功能前，AI 应优先检查：

项目已有的相关实现。

公共函数、组件和基础设施。

当前项目已经引入的依赖。

模块边界与复用可行性。

只有现有实现无法合理满足需求时，才新增组件。

### 12.2 复用限制
不强制为了复用而复用。

判断依据包括：

是否属于同一种业务职责。

是否具有一致的变化原因。

是否破坏模块边界。

是否增加不必要的耦合。

是否真实降低维护成本。

### 12.3 版本策略
V0.1：

通过 Skill 约束开发前的代码检查和复用决策。

V0.3：

增加组件索引、代码扫描和复用候选查询。

V1.0：

完善多语言分析和检测报告。

## 13. 工程检查与质量保障
### 13.1 Skill 约束
用于指导 Agent：

遵守技术选型。

复用已有组件。

保持模块职责。

维护项目上下文。

运行必要测试。

如实记录结果。

Skill 不能独立保证所有规则必然执行。

### 13.2 Go CLI 校验
用于检查：

项目配置

Schema 版本

技术兼容性

规则安装完整性

交接文档状态

Git 检查点信息

后续版本支持的架构依赖

### 13.3 CI 校验
后续将使用 CI 执行：

代码格式检查

静态分析

单元测试

架构依赖规则

必要项目配置检查

必须区分：

检查通过

检查失败

检查未执行

检查工具不可用

不能因无法执行检查而宣称通过。

## 14. 开发路线图
V0.1 — Foundation & Continuity
目标：实现可使用的跨平台基础技能包，并验证跨 AI 交接。

主要功能：

Go CLI

YAML 配置系统

架构选择

开发者确认

Agent Skills 安装

三平台适配

项目介绍和进度文档

AI 交接文档

检查点保存与恢复

基础配置与上下文检查

自动化测试

关键验收：

使用 Agent A 创建并开发项目，保存检查点后切换 Agent B；Agent B 在没有之前聊天记录的情况下，能够根据项目文档与代码继续未完成任务。

V0.2 — Architecture Validation
目标：增强工程规范的可执行性。

主要功能：

架构依赖检查

Go 包依赖分析

外部语言校验器适配

Git 变更信息校验

更可靠的交接文档新鲜度检查

CI 质量门禁

V0.3 — Code Intelligence
目标：提高 AI 对已有项目的理解能力。

主要功能：

现有项目扫描

代码符号索引

复用候选查询

重复实现提示

任务级上下文增强

历史工程问题基线

V1.0 — Stable Release
目标：提供稳定、可维护的开源全栈开发技能包。

主要功能：

完整核心 Skills

多语言、多框架支持

稳定的 CLI 和配置格式

完善的架构校验

代码复用分析

多 Agent 项目交接

1–5 人团队协作增强

跨平台发行包

完整使用与贡献文档

SoloWeave 应尽早公开源代码，V1.0 表示稳定版本，而不是首次开源。

## 15. V0.1 详细开发任务
阶段 A — Go 工程基础

| ID | 工作 |
| --- | --- |
| SW-001 | 创建 Go Module 和 Git 仓库 |
| SW-002 | 实现 Cobra CLI |
| SW-003 | 统一命令、错误类型和退出码 |
| SW-004 | 设计项目配置 Schema |
| SW-005 | 实现 YAML 读取、写入与校验 |
| SW-006 | 编写基础单元测试 |

验收：

Go 工程可以构建。

CLI 帮助和版本命令正常。

配置可以读写。

错误配置能够报告具体问题。

基础测试通过。

阶段 B — Project Continuity

| ID | 工作 |
| --- | --- |
| SW-007 | 定义 Project Context 文件规范 |
| SW-008 | 实现项目上下文初始化 |
| SW-009 | 实现 context show |
| SW-010 | 实现 context checkpoint |
| SW-011 | 实现 context resume |
| SW-012 | 实现 context check |
| SW-013 | 记录 Git 状态并检测过期检查点 |
| SW-014 | 编写上下文和交接测试 |

验收：

能生成项目介绍和进度模板。

可以保存带有 Git 信息的检查点。

可以查看当前恢复摘要。

可以发现明显缺失或过期的交接信息。

无需 AI API 即可完成基础操作。

阶段 C — 架构与技术选型

| ID | 工作 |
| --- | --- |
| SW-015 | 建立技术栈与架构目录 |
| SW-016 | 实现技术兼容性校验 |
| SW-017 | 实现交互式初始化 |
| SW-018 | 生成 Draft 项目配置 |
| SW-019 | 实现本地确认 |
| SW-020 | 生成初始化 ADR |

验收：

能选择完整技术方案。

不合法组合被拒绝。

生成项目配置和 ADR。

关键架构决策未经确认不视为生效。

阶段 D — Agent 平台适配

| ID | 工作 |
| --- | --- |
| SW-021 | 设计 Adapter 接口 |
| SW-022 | 实现 Codex 适配 |
| SW-023 | 实现 Claude Code 适配 |
| SW-024 | 实现 Cursor 适配 |
| SW-025 | 实现无覆盖安装及版本检查 |

验收：

三个平台能够安装适用规则。

重复安装不损坏文件。

原有项目规则保持完整。

安装状态可以检查。

阶段 E — 核心 Skills

| ID | 工作 |
| --- | --- |
| SW-026 | 编写 soloweave 主 Skill |
| SW-027 | 编写 project-setup |
| SW-028 | 编写 feature-workflow |
| SW-029 | 编写 project-continuity |
| SW-030 | 编写 quality-review |
| SW-031 | 建立 Agent 行为测试 |

验收：

AI 能读取项目配置。

AI 能读取并更新开发进度。

AI 在开发前检查现有实现。

AI 在任务结束时生成交接信息。

不同 Agent 能恢复既有任务。

阶段 F — 检查与发布

| ID | 工作 |
| --- | --- |
| SW-032 | 实现基础 check |
| SW-033 | 实现 doctor |
| SW-034 | 完成端到端测试 |
| SW-035 | 完成跨平台构建 |
| SW-036 | 编写 README 和快速开始文档 |
| SW-037 | 建立 CI |
| SW-038 | 发布 V0.1 预览版 |

验收：

完成项目初始化到 Skill 安装的完整流程。

项目上下文能够保存和恢复。

不同 AI 平台之间能够完成交接测试。

核心 Go 测试通过。

文档命令与程序行为一致。

## 16. 开源治理
### 16.1 开源策略
SoloWeave 从早期开发阶段公开仓库。

采用 Git 管理代码、文档和版本。

首次公开发布前选择适当的开源许可证，例如 MIT 或 Apache-2.0，并检查名称及分发标识的潜在冲突。

### 16.2 基础开源文件
需要包含：

README.md

LICENSE

CONTRIBUTING.md

CODE_OF_CONDUCT.md

SECURITY.md

CHANGELOG.md

GitHub Issue 模板

Pull Request 模板

### 16.3 贡献方向
社区可以贡献：

不同语言和框架的技术目录。

新的 Agent 平台适配器。

全栈开发 Skills。

项目架构模板。

测试与安全规则。

静态分析工具适配。

项目连续性改进。

### 16.4 兼容原则
核心工程配置、上下文文档及检查协议应尽量保持平台无关。

新增 Agent 平台时，不应要求重写项目的核心配置。

## 17. 首次开发计划
正式开发从以下步骤开始：

创建 SoloWeave Git 仓库。

确定 Go Module 地址与许可证。

建立 README、AGENTS.md 和基础开发文档。

实现 Go CLI 与项目配置。

实现 Project Continuity 的文件结构与基础命令。

实现架构选型和项目初始化。

实现 Codex、Claude Code、Cursor 适配器。

编写核心 Skills。

验证跨 AI 交接流程。

发布 V0.1 预览版。

最初开发时，不必创建全部目标目录。

优先保证已有模块设计清楚、代码复用合理、测试结果可靠。

SoloWeave 自身也应使用 SoloWeave 的项目上下文规范，从第一个开发阶段开始保存进度和交接记录。

## 18. V0.1 最终验收场景
假设开发者准备创建一个完整的 Web 应用。

第一阶段

使用 Codex 初始化项目，选择前后端技术栈，并确认项目架构。

第二阶段

Codex 开发注册、登录和用户管理功能，记录重要代码变更、实际测试结果和当前进度。

第三阶段

开发者需要切换账号或改用 Claude Code。

当前 Agent 保存检查点和交接文档。

第四阶段

Claude Code 从同一项目仓库读取 SoloWeave 的工程配置、开发状态和交接信息。

它核对源代码及 Git 状态后，确认哪些功能已经完成、当前任务是什么、下一步需要做什么。

第五阶段

Claude Code 按照项目既定架构继续开发，而不要求开发者重新完整介绍项目。

第六阶段

新的修改完成后，Claude Code 更新开发进度和交接信息，为下一次切换做好准备。

V0.1 核心验收条件
新 Agent 不需要原聊天记录。

能识别项目目标和技术架构。

能识别已完成与未完成任务。

能定位最近修改的相关文件。

能发现明显过期的交接信息。

不会把未执行的测试声明为通过。

不会擅自更改批准的架构。

能在规定工作流内继续开发。

开发完成后能够更新交接资料。

## 19. 最终产品定义
SoloWeave 是一套面向个人开发者与 1–5 人微型团队的开源 AI 全栈工程工作流与技能包。

它通过工程规范、项目上下文持久化和标准化开发流程，让 Codex、Claude Code、Cursor 等不同 AI 编程工具能够在统一的项目约束下持续协作。

它不替代开发者的判断，不绑定特定 AI 账号，也不把项目的重要知识局限在聊天记录中。

SoloWeave 的最终目标，是让项目可以持续推进，而不依赖某个特定 AI 的上下文。

文档状态：v0.2 开发基准版
