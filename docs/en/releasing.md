# Releasing Skills

[简体中文](../releasing.md)

Deliver five Skill ZIPs and SHA256SUMS. Users can install from GitHub without building an executable.

1. Review changes and client evidence. Bump each changed Skill's numeric X.Y.Z version, preserve slugs, and update bilingual docs and changelogs.
2. Run check-skills.ps1, check-docs.ps1, test-tooling.ps1, and package-skillhub.ps1. Inspect the five packages listed in dist/skills/SHA256SUMS, including references. Publish only files in the manifest; older packages may remain in the output directory.
3. On explicit instruction, commit in Chinese and push. Verify CI for that commit; local success is not remote CI success.
4. Install from the public URL in an isolated project and compare names, versions, references, and bytes. Unpushed changes are not publicly available.
5. On explicit publication instruction, create a GitHub Release and optionally attach the five ZIPs and checksum file. The Release identifies a bundle snapshot; Skill versions remain independent.
6. Check and update all five SkillHub slugs. Review GitHub import or upload individual ZIPs. Check versions, Chinese summaries, content, and references before submitting.
7. After approval, inspect public pages and downloads and update publication status. Submission, approval, and public availability are distinct.

CI artifacts are candidate downloads, not public Releases. Follow the [official installer documentation](https://github.com/vercel-labs/skills); inspect existing files and preserve user changes before any update.

## Free listings on Agensi

1. Obtain explicit authorization for platform publication. Build each complete Skill directory with SKILL.md and its required references; the existing packaging script produces that structure. Check every resource against skills/ and record ZIP hashes.
2. In the creator dashboard, inspect existing listings to avoid duplicates. Upload one ZIP per listing and select **Free / No charge**. For this project, use English listing text while preserving the original Chinese source package and stating its language and source version.
3. Review or clear generated metadata. Remove invented commands, unsupported compatibility claims and demonstrations that were not run. Declare actual file/terminal permissions and conditional network use; include the MIT source link and real verification limits.
4. Preview the price, description and file tree before submitting. Keep submission, review approval, public availability and real agent execution as separate results. Agensi may display a platform package version such as v1.0; it does not replace the source version stated in the description.
5. After approval, check public English text, Free pricing and required reference files. Record actual status in the [delivery snapshot](../release-status.md). A public page is not a test of the content loaded through MCP.

Format and workflow references: [creator checklist](https://www.agensi.io/learn/skill-md-creator-checklist), [submission guide](https://www.agensi.io/learn/how-to-sell-skills-on-agensi). Dashboard behavior is authoritative when documentation differs. In the 2026-10-18 free submissions, no payout connection or payment was needed.
