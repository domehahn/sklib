package spec_test

import (
	"testing"

	"github.com/domehahn/sklib/spec"
)

func TestNormalizePlatform_Canonical(t *testing.T) {
	cases := []struct {
		in   string
		want spec.Platform
	}{
		{"codex", spec.PlatformCodex},
		{"claude-code", spec.PlatformClaudeCode},
		{"gitlab-duo", spec.PlatformGitLabDuo},
		{"github-copilot", spec.PlatformGitHubCopilot},
		{"cursor", spec.PlatformCursor},
		{"windsurf", spec.PlatformWindsurf},
		{"openhands", spec.PlatformOpenHands},
		{"opencode", spec.PlatformOpenCode},
		{"ollama", spec.PlatformOllama},
		{"generic", spec.PlatformGeneric},
		{"all", spec.PlatformAll},
	}
	for _, tc := range cases {
		got, err := spec.NormalizePlatform(tc.in)
		if err != nil {
			t.Errorf("NormalizePlatform(%q) unexpected error: %v", tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("NormalizePlatform(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestNormalizePlatform_Aliases(t *testing.T) {
	cases := []struct {
		in   string
		want spec.Platform
	}{
		{"gitlab", spec.PlatformGitLabDuo},
		{"duo", spec.PlatformGitLabDuo},
		{"github", spec.PlatformGitHubCopilot},
		{"copilot", spec.PlatformGitHubCopilot},
		{"claude", spec.PlatformClaudeCode},
		{"open-code", spec.PlatformOpenCode},
		{"open-hands", spec.PlatformOpenHands},
		{"CLAUDE", spec.PlatformClaudeCode},
		{"  cursor  ", spec.PlatformCursor},
	}
	for _, tc := range cases {
		got, err := spec.NormalizePlatform(tc.in)
		if err != nil {
			t.Errorf("NormalizePlatform(%q) unexpected error: %v", tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("NormalizePlatform(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestNormalizePlatform_Unknown(t *testing.T) {
	_, err := spec.NormalizePlatform("vscode")
	if err == nil {
		t.Error("expected error for unknown platform, got nil")
	}
}

func TestNormalizePlatform_Empty(t *testing.T) {
	_, err := spec.NormalizePlatform("")
	if err == nil {
		t.Error("expected error for empty platform, got nil")
	}
}

func TestIsKnownPlatform(t *testing.T) {
	if !spec.IsKnownPlatform("claude-code") {
		t.Error("claude-code should be known")
	}
	if !spec.IsKnownPlatform("claude") {
		t.Error("claude alias should be known")
	}
	if spec.IsKnownPlatform("notaplatform") {
		t.Error("notaplatform should not be known")
	}
}

func TestExpandAllPlatforms(t *testing.T) {
	input := []spec.Platform{spec.PlatformAll}
	out := spec.ExpandAllPlatforms(input)
	if len(out) == 0 {
		t.Fatal("ExpandAllPlatforms returned empty slice")
	}
	for _, p := range out {
		if p == spec.PlatformAll {
			t.Error("ExpandAllPlatforms result must not contain 'all'")
		}
	}
}

func TestExpandAllPlatforms_Dedup(t *testing.T) {
	input := []spec.Platform{spec.PlatformCodex, spec.PlatformAll, spec.PlatformCodex}
	out := spec.ExpandAllPlatforms(input)
	seen := make(map[spec.Platform]int)
	for _, p := range out {
		seen[p]++
	}
	for p, count := range seen {
		if count > 1 {
			t.Errorf("duplicate platform %q in ExpandAllPlatforms result", p)
		}
	}
}

func TestNormalizePlatforms_Dedup(t *testing.T) {
	out, err := spec.NormalizePlatforms([]string{"codex", "claude", "claude-code", "codex"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// claude and claude-code both resolve to claude-code; codex appears twice
	// expected: [codex, claude-code]
	if len(out) != 2 {
		t.Errorf("expected 2 unique platforms, got %d: %v", len(out), out)
	}
}

func TestPlatformStrings(t *testing.T) {
	in := []spec.Platform{spec.PlatformCodex, spec.PlatformCursor}
	got := spec.PlatformStrings(in)
	if got[0] != "codex" || got[1] != "cursor" {
		t.Errorf("unexpected PlatformStrings result: %v", got)
	}
}
