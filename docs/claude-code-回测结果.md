# Claude Code 回测结果

测试日期：2026-10-11（Asia/Shanghai）。本轮在本机完成，使用真实 Claude Code CLI，不是模拟会话。

## 环境与安装

- 系统：Windows 11 专业版 64 位（10.0.26200）；Claude Code CLI：2.1.218。CLI 使用 `--model sonnet`，会话界面实际显示模型 `glm-5.3`。
- 启动方式：在各隔离目录运行官方 `claude` CLI；场景一使用交互终端（TTY），场景二、三使用 `claude -p --output-format json`，场景三第二会话为全新会话。原始 JSONL 会话证据已保存到 `docs/claude-code-evidence/2026-10-11/`。
- SoloWeave 工作目录：`C:\Users\Administrator\Desktop\dmx\SoloWeave`；分支 `master`；HEAD `fd6e04a77e32bdff0d121fda753be4231a6f1333`。测试前后 Git 状态一致，均为既有 16 个已修改文件和 2 个未跟踪交接/报告文件；证据见 `git-before-evidence.txt`、`git-after-evidence.txt`。回测没有改写根目录 Skill 源码。
- 隔离 Skill 安装：复制根目录 `skills/ ` 六个完整目录到各测试目录的 `.claude/skills/`。场景一目录为 `C:\Users\Administrator\AppData\Local\Temp\soloweave-claude-retest-20261011-144413-s1b`；场景二为同名 `-s2b`；场景三为 `-s3`。六个版本均核对为：`soloweave 0.2.2`、`project-setup 0.3.1`、`feature-workflow 0.2.2`、`project-continuity 0.2.1`、`quality-review 0.2.0`、`systematic-debugging 0.1.0`。场景一、二的 `project-setup/references/selection-guide.md` 均存在。
- 客户端可观察到 Skill 事件：场景一终端显示 `Skill(soloweave) Successfully loaded skill`、`Skill(project-setup) Successfully loaded skill`、`Skill(feature-workflow) Successfully loaded skill`；场景三 JSONL 也保留了相关事件。证据文件：`s1-session.jsonl`、`s3-first-session.jsonl`、`s3-second-session.jsonl`。场景一另保存了 Edge headless 浏览器截图 `browser-screenshot.png`。

## 场景一：空目录待办应用

- 初始目录只有 `.claude/skills/`（六个 Skill 及参考文件），没有应用代码、README 或 SoloWeave 上下文；完整清单见 `s1b-initial-files.txt`。
- 原始提示词：

  > 请在当前目录开发一个简单的待办事项应用：可以新增、完成和删除任务，关闭后再次打开仍能看到任务。界面保持简洁。请根据目录现状开始处理。

- 首轮观察：Claude Code 先加载 `soloweave`，读取目录后明确判断“目录为空，属于新建项目”，再加载 `project-setup`。它在写应用文件前询问应用形态，并列出单文件网页、CLI、网页+后端、桌面应用等候选。选择“网页 + 后端”后又询问后端栈，列出 Node 零依赖、Node+Express、Python+Flask 等候选；选择 Node 零依赖后，先写 `PROJECT.md` 与 `ADR-0001.md`，再加载 `feature-workflow` 开发。关键选型确认前没有写应用代码。
- 实际新增文件：`server.js`、`public/index.html`、`test-api.mjs`、运行后生成的 `data/todos.db`，以及 `.soloweave/context/{PROJECT,STATUS,HANDOFF,CHANGES}.md` 和 `.soloweave/decisions/ADR-0001.md`。完整清单见 `s1b-final-files.txt`。
- 实际验证：
  - `node --version` → v24.18.0；`npm --version` → 11.16.0；`python --version` → 3.12.10。
  - 首次中文 `curl` 组合命令因 Windows shell 引号退出码 2；Claude Code 没有把它算作通过，改用 Node UTF-8 冒烟脚本。
  - `node test-api.mjs` 实际运行通过：8 项断言 ALL PASS，包含新增、列表、完成、删除、空文本 400、404、中文无损。
  - 停止并重启服务后，实际运行持久化检查输出 `PERSIST PASS`；随后删除测试任务，剩余列表为 `[]`。
  - 使用 Edge headless 实际加载页面并查看截图；页面标题、输入框、添加按钮和空状态正常。通过同源浏览器测试页实际执行新增、勾选完成、iframe 重载（模拟关闭重开）和删除，结果为 `{added:true, persistedDone:true, deleted:true}`。Node 24 的 `node:sqlite` 实验性警告已观察到，不影响上述结果。
- 结论：**通过关键路由、浏览器功能和确认边界**。真实客户端在空目录先进入项目设置，先询问影响项目形态的选择，确认后才实现；API、浏览器新增/完成/删除和重载持久化均通过。

## 场景二：分维度技术推荐

- 隔离目录初始同样只有六个 Skill，无应用文件；清单见 `s2b-initial-files.txt`。
- 原始提示词：

  > 请为一个前后端分离的待办应用讨论技术选型。先了解我的需求，再按前端、后端、数据、架构和交付等适用维度给可组合的候选与推荐；暂时不要写代码。

- 首轮回复先询问规模、平台、功能边界、离线/同步、团队背景、部署预算和非功能要求，没有写代码。为让推荐矩阵基于事实，补充了：1–5 人团队、桌面和手机浏览器、MVP 增删改查与持久化、简单账号、JS/TS、单台 Linux 云服务器、预算有限。
- 随后回复按前端、后端、数据、认证、架构/API、交付分别给出候选、收益、约束、推荐和平级替代。例如 React/Vue/Svelte；Node Fastify/Hono/Express；SQLite/PostgreSQL；REST/tRPC；Docker Compose/systemd。明确推荐仍是候选，没有写入批准记录，并指出版本/兼容性需要核对；网络搜索请求未获授权，未把不确定版本当作已验证事实。
- 结论：**通过**。推荐按适用维度拆分且可组合，解释了取舍和兼容性，没有直接写代码或把建议记成批准决定。

## 场景三：受阻后的新会话接手

- 阻塞测试目录：`C:\Users\Administrator\AppData\Local\Temp\soloweave-claude-retest-20261011-144413-s3`。预置真实未完成项目：`package.json` 的 build 脚本为 `command-that-does-not-exist-soloweave`，`src/index.js` 仅输出 unfinished；README 明确要求记录阻塞、不得改脚本或绕过。
- 第一会话提示词：

  > 请开始处理当前项目：先读现有文件和 SoloWeave Skills，运行现有构建检查；如果现有构建因缺少工具或依赖失败，请把真实错误写入 .soloweave/context/STATUS.md 和 HANDOFF.md，保留未完成状态并停止修复，不要修改构建脚本或绕过阻塞。

  Claude Code 加载 `soloweave`/交接流程，创建并填写 `.soloweave/context/PROJECT.md`、`STATUS.md`、`HANDOFF.md`、`CHANGES.md`。其会话权限拦截了 `npm run build` 及其他执行命令，因此它把构建标为“未运行”，没有伪造运行时报错，也没有改代码或构建脚本。
- 为取得实际阻塞证据，随后在测试目录独立运行同一命令：`npm run build` 输出 `command-that-does-not-exist-soloweave is not recognized...`，退出码 1。该命令由 Codex 外部核对，报告中与 Claude 会话内“权限拦截、未运行”分开记录。
- 第二会话为不带旧聊天记录的新 Claude Code 会话，唯一提示词：

  > 请继续这个项目

  它先读取并核对 `STATUS.md`、`HANDOFF.md` 与实际文件，重新尝试构建仍被权限模式拦截，保持“未完成、不可运行”结论，明确代码和构建脚本零改动，并更新交接记录。没有把项目称为可运行。
- 结论：**通过交接行为**；真实阻塞、权限限制、未运行检查和下一步均被保留。构建命令的外部实跑退出码为 1；Claude 会话自身未获执行权限，故没有把外部结果冒充其会话输出。

## 总结

- 已观察到的符合预期行为：空目录先进入 `project-setup`；关键形态和后端选择确认前不写应用代码；分维度提供可组合技术候选；受阻时写入事实、验证状态和下一步；新会话先读交接且不宣称未验证项目可运行。
- 已观察到的限制：Claude Code 的 Bash 权限模式会拦截部分命令，需要逐项批准；Windows shell 传入中文的 curl 命令出现编码/引号问题，改用 Node UTF-8 脚本后通过。浏览器测试使用 Edge headless 和临时同源代理；场景二未做网络版本核验。
- 尚未验证：Claude Code 之外的 Cursor 真实客户端、公开 SkillHub 更新、人工逐步点击的完整浏览器验收，以及不同权限设置下的自动调用行为。
- 建议（观察与推测分开）：观察事实是新版六个 Skill 在本机 Claude Code 会话中被实际加载，三场景均按预期路由/交接；推测性建议是把“Windows 下中文 curl 可能受 shell 编码影响”和 Bash 权限批准作为验收环境说明，并在后续加入一次浏览器手动检查。
