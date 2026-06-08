# sklib

`sklib` is the shared Go library for AI Agent Skill tooling.

It provides canonical data models, parsing helpers, validation primitives, and registry API DTOs used across the skill ecosystem tools:

| Tool | Role |
| --- | --- |
| **skcr** | Creates skill skeletons and renders project/platform structures |
| **skpm** | Generic skill package manager and lifecycle manager |
| **SkillForge** | Concrete skill registry server |
| **skforge** | SkillForge-native CLI/admin tool |

`sklib` prevents logic drift between these tools by providing a single source of truth for all shared types and validation rules.

---

## What sklib contains

| Package | Purpose |
| --- | --- |
| `spec` | Canonical data models: `Skill`, `AgentSkillsManifest`, `AgentSkillsLock`, `PackageManifest`, platform constants, SemVer helpers |
| `ref` | Skill reference parsing and formatting (`namespace/name@constraint`) |
| `validate` | Composable validation result types and validators |
| `registryapi` | Registry HTTP API DTOs (request/response structs, error codes) |
| `packageio` | Package artifact layout constants, checksum parsing, manifest I/O |
| `specio` | YAML/JSON read-write helpers for spec files |

---

## What sklib intentionally does not contain

```text
sklib does not scaffold skills.
sklib does not package skills.
sklib does not publish skills.
sklib does not install skills.
sklib does not implement HTTP clients.
sklib does not implement a registry server.
sklib does not implement full package manager workflows.
sklib does not import skcr, skpm, SkillForge, or skforge.
sklib only provides shared models and validation primitives.
```

Those responsibilities belong to `skcr`, `skpm`, `SkillForge`, and `skforge` respectively.

---

## Install

```sh
go get github.com/domehahn/sklib
```

---

## Example imports

```go
import "github.com/domehahn/sklib/spec"
import "github.com/domehahn/sklib/ref"
import "github.com/domehahn/sklib/validate"
import "github.com/domehahn/sklib/registryapi"
import "github.com/domehahn/sklib/packageio"
import "github.com/domehahn/sklib/specio"
```

---

## Example: define and validate a skill

```go
import (
    "github.com/domehahn/sklib/spec"
    "github.com/domehahn/sklib/validate"
)

skill := spec.Skill{
    Name:        "gitlab-policy-reviewer",
    Namespace:   "myorg",
    Version:     "1.5.0",
    Description: "Reviews GitLab security policies.",
    Entrypoint:  "SKILL.md",
    CompatibleWith: []spec.Platform{
        spec.PlatformClaudeCode,
        spec.PlatformGitLabDuo,
    },
}

result := validate.ValidateSkillMetadata(skill, validate.Options{Strict: true})
if !result.Valid {
    for _, f := range result.Findings {
        fmt.Printf("[%s] %s: %s\n", f.Severity, f.Field, f.Message)
    }
}
```

---

## Example: parse a skill reference

```go
import "github.com/domehahn/sklib/ref"

r, err := ref.ParseSkillRef("myorg/gitlab-policy-reviewer@^1.5.0")
if err != nil {
    log.Fatal(err)
}
fmt.Println(r.Namespace, r.Name, r.Constraint)
// myorg  gitlab-policy-reviewer  ^1.5.0
```

---

## Example: normalize platforms

```go
import "github.com/domehahn/sklib/spec"

platforms, err := spec.NormalizePlatforms([]string{"claude", "gitlab", "all"})
// "claude" -> PlatformClaudeCode
// "gitlab" -> PlatformGitLabDuo
// "all"    -> expanded to all concrete platforms

expanded := spec.ExpandAllPlatforms(platforms)
```

---

## Example: read a skill.yaml

```go
import "github.com/domehahn/sklib/specio"

skill, err := specio.ReadSkillYAML("path/to/skill.yaml")
if err != nil {
    log.Fatal(err)
}
fmt.Println(skill.Name, skill.Version)
```

---

## Example: registry API DTOs

```go
import "github.com/domehahn/sklib/registryapi"

caps := registryapi.FullCapabilities("v1")
// caps.Features.Resolve == true
// caps.Features.Publish == true

errResp := registryapi.NewErrorResponse(registryapi.CodeNotFound, "skill not found")
```

---

## Platform constants

```go
spec.PlatformCodex         // "codex"
spec.PlatformClaudeCode    // "claude-code"
spec.PlatformGitLabDuo     // "gitlab-duo"
spec.PlatformGitHubCopilot // "github-copilot"
spec.PlatformCursor        // "cursor"
spec.PlatformWindsurf      // "windsurf"
spec.PlatformOpenHands     // "openhands"
spec.PlatformOpenCode      // "opencode"
spec.PlatformOllama        // "ollama"
spec.PlatformGeneric       // "generic"
spec.PlatformAll           // "all" (expands to all concrete platforms)
```

Aliases accepted by `NormalizePlatform`: `claude`, `gitlab`, `duo`, `github`, `copilot`, `open-code`, `open-hands`.

---

## SemVer helpers

```go
spec.IsSemVer("1.2.3-alpha")        // true
spec.IsStableSemVer("1.2.3")        // true
spec.IsStableSemVer("1.2.3-alpha")  // false
v, err := spec.NormalizeVersion("v1.2.3") // "1.2.3", nil
spec.CompareVersions("1.2.3", "1.3.0")   // -1
```

---

## Relationship to skillspec

[`skillspec`](https://github.com/domehahn/skillspec) defines the vendor-neutral specification: JSON Schemas, OpenAPI, and package-layout documentation.

`sklib` is the Go implementation of that specification. Types, validation rules, and field names in `sklib` are derived from and must stay consistent with `skillspec`.

---

## Module

```text
github.com/domehahn/sklib
```

Go 1.21+ required.
