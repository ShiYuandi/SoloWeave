# SoloWeave v{{VERSION}} Windows x64 安装包

本压缩包已包含可运行的 `soloweave.exe`、中英文使用说明和 MIT 许可证。**无需安装 Go，也无需自行构建。**

此程序是旧版 CLI 快照，内置五个旧版 Skill；新版六个 Skills 位于 [SoloWeave 仓库](https://github.com/ShiYuandi/SoloWeave) 的 `skills/`，不会自动进入本程序。

## 开始使用

1. 解压 ZIP，将所在目录加入当前终端的 `PATH`，或在下面的命令中使用 `soloweave.exe` 的完整路径。
2. 在目标项目目录运行：

```powershell
soloweave init
```

审阅生成的 `.soloweave/project.yaml`；技术方案由开发者确认后再运行：

```powershell
soloweave approve
soloweave install --agents codex,claude,cursor --dry-run
soloweave install --agents codex,claude,cursor
soloweave doctor
soloweave context checkpoint --task "初始化项目" --summary "已确认配置并安装 Skills" --next "开始开发"
soloweave check
```

`--agents` 可只填写你使用的 AI 平台。`install --dry-run` 只预览，正式 `install` 才写入五个 Skills 和项目规则。`check` 需要先记录检查点。任务交接时再次使用 `soloweave context checkpoint`；新会话使用 `soloweave context resume`。没有运行测试时，不要填写 `--verification`。

只想使用开发规范、不需要 CLI 时，可以单独从 SkillHub 安装五个 Skills。两种方式与各命令的完整说明见 [GitHub README](https://github.com/ShiYuandi/SoloWeave#readme) 和[纯 Skills 指南](https://github.com/ShiYuandi/SoloWeave/blob/master/docs/skill-only-workflow.md)。本文件随当前安装包的版本生成；最新公开版本以 [GitHub Releases](https://github.com/ShiYuandi/SoloWeave/releases) 为准。
