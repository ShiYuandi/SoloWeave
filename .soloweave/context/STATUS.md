# 当前状态

核对日期：2026-10-11。

## 已完成的本地工作

- 按开发者确认的八阶段计划先写 `docs/go-retirement-plan.md` 和 ADR-0003，再实施纯 Skills 迁移；保留原有未提交的路由与文档修订。
- 移除 Go 源码、Module、旧嵌入 Skills、程序构建与发布脚本；清理本地 Go 工具链与旧 CLI 构建产物。
- 旧 `project.yaml`、`approval.json`、`checkpoint.json` 和使用说明已归档；Markdown 上下文与 ADR 保留。
- 六个本地 Skill 候选：入口和功能开发 0.2.3、项目设置 0.3.2、项目连续性 0.2.2、质量审查 0.2.1、故障排查 0.1.0。
- 新增维护者校验与负例测试，打包包含参考资源和 SHA256SUMS；CI 已改为纯 Skills 检查。
- 同步中英文 README、使用、贡献、目录、配置、交接、迁移和发布说明；历史设计标为历史。

## 实际运行的验证

- `scripts/check-skills.ps1`：六个 Skill 的单行元数据、唯一 slug、版本和资源边界检查通过；此脚本不是通用 YAML 解析器。
- `scripts/test-tooling.ps1`：15 项通过，包括无效配置负例、禁止包输出进入源码目录、六包校验和、全部资源字节一致及重复打包一致。
- `scripts/package-skillhub.ps1`：生成六包及 SHA256SUMS，位于忽略目录 `dist/skills/`。
- `scripts/check-docs.ps1` 与 `git diff --check`：通过；文档脚本只检查相对文件链接，不检查网页和标题锚点。
- 缓存的标准安装器 skills@1.7.1 发现六个 Skill；六个新建临时项目分别只安装一个 Skill 到 Codex、Claude Code、Cursor 的项目目录，安装文件 SHA256 与源码一致。这不证明自动触发或已有文件冲突处理。
- 三个归档机器文件与 HEAD 对应内容核对通过（换行归一化）。
- 已导入跨设备 Claude Code CLI 报告及 17 个证据文件；实际模型为 glm-5.3。原始调用支持旧修订版的新项目路由，受阻接手未观察到单独 continuity 调用。本机未重跑测试应用。
- 后续真实 Codex CLI 回测完成 12 个独立临时会话，当前候选版六项独立最小行为、全套空目录路由、受阻双会话接手与获准修复通过。实际读到各本地 Skill；外部核对复用实现、原测试哈希、失败与修复结果、未确认时无应用代码、受阻时未伪造构建工具。结果见 docs/codex-behavior-results.md；新增可复跑夹具和 CLI 运行脚本，未更改 Skill 行为或版本。

## 受阻与未执行

- 首次 Codex CLI 读取拒绝作为历史保留，后续回测已经成功；未关闭沙箱或执行规则，不能据此断言旧失败的单一根因。
- Cursor、Anthropic 模型及更多项目类型未在本轮实测；一次最小场景通过不能保证所有任务总会自动触发。
- skill-creator 辅助 quick_validate.py 因缺少 PyYAML 未成功；不影响仓库自有脚本已运行的结果。
- SkillHub 六项审核已提交；核对时交接 0.2.2 已发布，其余五项安全审核中，待审版本公开下载尚未验收。见 docs/release-status.md。

## 下一步

当前本地行为验收和本轮授权的线上提交已完成。后续核对 SkillHub 审核及公开下载；Cursor、Anthropic 模型和更多场景作为后续独立回归。不要重新全局安装用户已删除的 Skills。

## 本轮线上操作授权

开发者已明确授权中文提交、推送，确认 GitHub CI 与公开安装后提交六项 SkillHub 审核。已执行：迁移 eb452f5、CI 修复 135e13c 推送至 master；135e13c 远端 CI 通过。15 项检查本来已通过，首次 CI 误用负例子进程退出码，修复明确成功退出码后复验成功。

公开 GitHub 安装使用 skills@1.7.1，在全新临时项目选择 Codex/Claude Code/Cursor，安装六项；.agents/skills 与 .claude/skills 共 14 个文件逐项 SHA256 与源码一致。没有全局安装。SkillHub 五项通过 ZIP 更新，故障排查通过 GitHub 导入新增；六项均已提交，项目设置含参考资料。截图在 dist/acceptance/skillhub-submitted-2026-10-11.jpg；实际快照见 docs/release-status.md。
