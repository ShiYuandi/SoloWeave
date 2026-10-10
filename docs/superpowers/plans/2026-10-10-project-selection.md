# SoloWeave 项目选型推荐实施计划

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [x]`) syntax for tracking.

**Goal:** 让 `project-setup` 从需求出发识别项目类型，按适用维度提出可组合、有依据且不受预设技术清单限制的选型建议，并经开发者确认后记录。

**Architecture:** `SKILL.md` 规定交互顺序和确认边界，按需读取的参考文件只提供判断问题、比较标准和非约束性示例。行为场景和使用文档说明预期与真实客户端验收；不修改 Go CLI。

**Tech Stack:** Agent Skills Markdown、PowerShell 打包脚本、现有中英文 Markdown 文档。

**Spec:** [项目选型与推荐重设计](../specs/2026-10-10-project-selection-design.md)

## Global Constraints

- 保留六个 Skill 的名称及 SkillHub slug；本轮只修改根目录 `skills/project-setup/` 的选型行为，不改旧版 Go CLI 的内置资源或机器批准文件。
- 类型覆盖 Web 前端、后端 API、全栈 Web、移动端、桌面端、CLI／自动化、数据／AI；允许混合和未列出的类型。类型用于发现问题，不限定技术。
- 参考文件里的路线与技术只是例子，不是白名单、默认技术栈、固定排序或数量上限。AI 可依据需求提出未列出的技术和组合。
- 已有项目以实际代码、依赖和有效 ADR 为起点；用户已给的选型不重复追问；关键新选型经明确确认才写成已批准 ADR。
- 保留工作区尚未提交的 `0.2.1` 交接修订。未经用户明确要求，不提交、不推送、不发布。
- 静态打包成功与真实 Agent 行为分开报告；没有运行过的验证不得称为通过。

## Review Focus

- 移动端加后端 API：应分别比较适用维度并检查接口和部署约束，不强制单仓库或多仓库；归 Task 3 的客户端场景。
- 只做数据分析而不使用模型：不应强迫选择模型服务商；归 Task 2 的参考规则与 Task 3 的场景。
- 用户提出参考文件外的技术：应正常评估并允许加入组合；归 Task 2 的参考规则与 Task 3 的场景。
- 当前版本或兼容性无法核实：应标记不确定性与待验证事项，不编造事实；归 Task 2 的参考规则与 Task 3 的场景。
- 用户已明确指定全部关键选型：应核对相容性并复述待确认组合，不重复问卷；归 Task 1 的主规则与 Task 3 的场景。

---

### Task 1：重写项目设置的行为契约

**Files:**
- Modify: `skills/project-setup/SKILL.md`

**Interfaces:**
- Consumes: 设计文档中“交互流程”“异常与边界”；仓库现有 `PROJECT.md`、ADR、CLI 兼容约定。
- Produces: Skill 触发描述、需求澄清 → 类型与维度识别 → 候选比较 → 组合检查 → 开发者确认 → 文档记录的流程；按需引用 `references/selection-guide.md`。

- [x] **Step 1:** 核对 `SKILL.md` 当前 frontmatter、版本、slug，以及其他五个 Skill 的版本；确定本次递增版本并保持 slug 不变。
- [x] **Step 2:** 修改触发描述和正文，明确类型用于选择问题维度，AI 根据需求自行生成候选，参考示例不得限制技术；保留已有项目沿用、待确认和 ADR 的边界。
- [x] **Step 3:** 检查两条输入路径：模糊“做一个待办应用”不得直接假定静态网页；“前后端分离且已指定技术”不得重复询问已给信息或仅提出绑定好的三套组合。
- [x] **Step 4:** 核对文件中没有强制固定技术栈、固定候选数量、未经确认的批准或纯 Skills 手工生成 `approval.json` 的指令。

### Task 2：提供按需读取的选型判断指南

**Files:**
- Create: `skills/project-setup/references/selection-guide.md`

**Interfaces:**
- Consumes: Task 1 的按需引用、设计文档的项目范围与比较标准。
- Produces: 按项目类型和决策维度组织的关键问题、取舍标准、典型兼容性检查、少量非约束性路线示例；供 Task 1 在需要选型时引用。

- [x] **Step 1:** 写共同判断问题：目标、平台、数据、离线、部署、团队经验、成本、维护与既有约束；说明仅追问会改变关键决定的缺口。
- [x] **Step 2:** 写 Web、移动端、桌面端、CLI／自动化、数据／AI 的适用维度与兼容性检查；覆盖混合类型，并在数据分析无模型时跳过模型维度。
- [x] **Step 3:** 为不同路线放少量示例，明确 AI 可增减候选、接受表外技术；对无法核实的版本和兼容性记录不确定性。
- [x] **Step 4:** 对照设计文档检查没有封闭列表、无条件默认栈、强制 Web 维度或把互不冲突的架构维度当成互斥方案。

### Task 3：扩展验收场景并验证安装产物

**Files:**
- Modify: `docs/agent-acceptance.md`
- Modify: `docs/en/agent-acceptance.md`
- Modify: `scripts/package-skillhub.ps1`

**Interfaces:**
- Consumes: Tasks 1–2 的 Skill 与参考文件。
- Produces: 可复现的客户端提示词和通过条件、真实结果栏位；包含 Skill 全部所需文件的本地 ZIP 与内容一致性记录。

- [x] **Step 1:** 在双语验收文档加入模糊待办、前后端分离、多类型混合、移动端、CLI、纯数据分析、表外技术、已有项目、预设全部选型、不兼容组合和无法核实版本的场景；每项写明应观察的关键行为。
- [x] **Step 2:** 修改 `scripts/package-skillhub.ps1`，在保留现有 frontmatter 校验及 ZIP 命名的前提下，把每个 Skill 目录中的所需文件按相对路径打包；`project-setup` 包须同时包含 `SKILL.md` 和 `references/selection-guide.md`。
- [x] **Step 3:** 运行 `pwsh -NoProfile -File scripts/package-skillhub.ps1`；预期退出码 0。核对六个 ZIP 的文件条目、逐文件字节与源文件一致，并核对 `SHA256SUMS`。
- [x] **Step 4:** 用标准安装器对本地仓库执行 `add . --list`，再在隔离目录安装 `project-setup`；预期安装后的 `SKILL.md` 和 `references/selection-guide.md` 均与源文件一致。
- [x] **Step 5:** 运行 `git diff --check`，预期无输出且退出码 0。
- [ ] **Step 6:** 在安装新版 Skill 的真实 Codex 新会话执行 Task 3 的核心场景，记录实际安装版本、原提示词、Agent 回复和项目文件。若本会话无法操作另一客户端，则标为“待真实客户端复测”，不得视为通过。

### Task 4：同步公开说明和交接记录

**Files:**
- Modify: `docs/skills.md`, `docs/en/skills.md`
- Modify: `docs/skill-only-workflow.md`, `docs/en/skill-only-workflow.md`
- Modify as needed: `README.md`, `docs/en/README.md`, `CHANGELOG.md`, `docs/en/CHANGELOG.md`
- Modify: `.soloweave/context/STATUS.md`, `.soloweave/context/HANDOFF.md`, `.soloweave/context/CHANGES.md`

**Interfaces:**
- Consumes: Tasks 1–3 的实际行为与验证结果。
- Produces: 中英文一致的用户说明、版本记录及可供下一会话接手的真实状态。

- [x] **Step 1:** 更新使用说明：先描述需求、确认项目类型与适用维度，再逐项选择或自定义技术，最后确认组合；明确参考例子不限制 AI。
- [x] **Step 2:** 按实际版本更新中英文变更记录；README 只在现有介绍需要修正时改动，不重复长篇指南。
- [x] **Step 3:** 检查改动 Markdown 的相对链接、`git diff --check` 和 `git status --short`；只写入实际运行的结果。
- [x] **Step 4:** 更新 STATUS、HANDOFF、CHANGES，分别记录完成项、未完成的客户端复测、真实验证与下一步；不将旧 `checkpoint.json` 当成本轮检查点。

## 执行与交付边界

按 Task 1→4 顺序实施。交接修订 `0.2.1` 与本次项目选型改动需在检查点中分别标识，避免把两者的验证结果混为一谈。完成本地实现后，若真实客户端复测仍待用户运行，交付应明确指出可安装与行为验收的区别。提交、推送和 SkillHub 更新等待用户另行明确要求。

## 实施记录

- 本地 Skill、参考指南、打包脚本及双语说明已按计划修改；SkillHub 和 GitHub 尚未更新。
- 本地六包及标准安装器隔离安装已核对；真实 Codex 新会话选型行为仍待复测，Task 3 Step 6 保持未完成。
- 仓库约定要求只有用户明确提出时才提交或推送；本轮未提交。
