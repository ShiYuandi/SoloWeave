# Releasing Skills

[简体中文](../releasing.md)

Deliver six Skill ZIPs and SHA256SUMS. Users can install from GitHub without building an executable.

1. Review changes and client evidence. Bump each changed Skill's numeric X.Y.Z version, preserve slugs, and update bilingual docs and changelogs.
2. Run check-skills.ps1, check-docs.ps1, test-tooling.ps1, and package-skillhub.ps1. Inspect the six packages listed in dist/skills/SHA256SUMS, including references.
3. On explicit instruction, commit in Chinese and push. Verify CI for that commit; local success is not remote CI success.
4. Install from the public URL in an isolated project and compare names, versions, references, and bytes. Unpushed changes are not publicly available.
5. On explicit publication instruction, create a GitHub Release and optionally attach the six ZIPs and checksum file. The Release identifies a bundle snapshot; Skill versions remain independent.
6. Check and update all six SkillHub slugs. Review GitHub import or upload individual ZIPs. Check versions, Chinese summaries, content, and references before submitting.
7. After approval, inspect public pages and downloads and update publication status. Submission, approval, and public availability are distinct.

CI artifacts are candidate downloads, not public Releases. Follow the [official installer documentation](https://github.com/vercel-labs/skills); inspect existing files and preserve user changes before any update.
