package validate_test

import (
	"testing"

	"github.com/domehahn/sklib/spec"
	"github.com/domehahn/sklib/validate"
)

// ─── Result tests ────────────────────────────────────────────────────────────

func TestResult_Finalize(t *testing.T) {
	var r validate.Result
	r.AddError("field", "something broke")
	r.Finalize()
	if r.Valid {
		t.Error("Result with errors should not be valid after Finalize")
	}
}

func TestResult_FinalizeNoErrors(t *testing.T) {
	var r validate.Result
	r.AddWarning("field", "just a warning")
	r.Finalize()
	if !r.Valid {
		t.Error("Result with only warnings should be valid after Finalize")
	}
}

func TestResult_HasErrors(t *testing.T) {
	var r validate.Result
	if r.HasErrors() {
		t.Error("empty result should have no errors")
	}
	r.AddWarning("", "w")
	if r.HasErrors() {
		t.Error("result with only warning should have no errors")
	}
	r.AddError("", "e")
	if !r.HasErrors() {
		t.Error("result with error should have errors")
	}
}

func TestFindingSeverity(t *testing.T) {
	var r validate.Result
	r.AddError("f1", "err msg")
	r.AddWarning("f2", "warn msg")
	r.AddInfo("f3", "info msg")
	if len(r.Findings) != 3 {
		t.Fatalf("expected 3 findings, got %d", len(r.Findings))
	}
	if r.Findings[0].Severity != validate.SeverityError {
		t.Errorf("expected error severity, got %q", r.Findings[0].Severity)
	}
	if r.Findings[1].Severity != validate.SeverityWarning {
		t.Errorf("expected warning severity, got %q", r.Findings[1].Severity)
	}
	if r.Findings[2].Severity != validate.SeverityInfo {
		t.Errorf("expected info severity, got %q", r.Findings[2].Severity)
	}
}

// ─── Skill metadata tests ─────────────────────────────────────────────────────

func TestValidateSkillMetadata_Valid(t *testing.T) {
	s := spec.Skill{
		Name:           "my-skill",
		Version:        "1.0.0",
		Description:    "Does things.",
		CompatibleWith: []spec.Platform{spec.PlatformCodex},
		Entrypoint:     "SKILL.md",
	}
	r := validate.ValidateSkillMetadata(s, validate.Options{Strict: true})
	if !r.Valid {
		t.Errorf("expected valid, got findings: %+v", r.Findings)
	}
}

func TestValidateSkillMetadata_StrictEmptyDescription(t *testing.T) {
	s := spec.Skill{
		Name:           "my-skill",
		Version:        "1.0.0",
		CompatibleWith: []spec.Platform{spec.PlatformCodex},
	}
	r := validate.ValidateSkillMetadata(s, validate.Options{Strict: true})
	if r.Valid {
		t.Error("expected invalid (missing description in strict mode)")
	}
}

func TestValidateSkillMetadata_PermissiveEmptyDescription(t *testing.T) {
	s := spec.Skill{Name: "my-skill", Version: "1.0.0"}
	r := validate.ValidateSkillMetadata(s, validate.Options{Strict: false})
	// should produce warnings, not errors
	if !r.Valid {
		t.Errorf("expected valid in permissive mode, findings: %+v", r.Findings)
	}
	var hasWarn bool
	for _, f := range r.Findings {
		if f.Severity == validate.SeverityWarning {
			hasWarn = true
		}
	}
	if !hasWarn {
		t.Error("expected at least one warning for missing description")
	}
}

// ─── Filesystem tests ─────────────────────────────────────────────────────────

func TestValidateNoAbsoluteLocalPaths(t *testing.T) {
	r := validate.ValidateNoAbsoluteLocalPaths([]string{"/etc/passwd", "../secret"})
	if r.Valid {
		t.Error("expected invalid for absolute/traversal paths")
	}
	if !r.HasErrors() {
		t.Error("expected errors for forbidden paths")
	}
}

func TestValidateNoAbsoluteLocalPaths_Valid(t *testing.T) {
	r := validate.ValidateNoAbsoluteLocalPaths([]string{"SKILL.md", "src/main.go"})
	if !r.Valid {
		t.Errorf("expected valid, findings: %+v", r.Findings)
	}
}

func TestValidateNoObviousSecrets(t *testing.T) {
	r := validate.ValidateNoObviousSecrets([]string{"config/.env", "certs/server.pem"})
	if len(r.Findings) == 0 {
		t.Error("expected findings for obvious secret paths")
	}
	// Should be warnings, not errors
	for _, f := range r.Findings {
		if f.Severity == validate.SeverityError {
			t.Errorf("expected warning, got error: %s", f.Message)
		}
	}
}

func TestValidateNoForbiddenFiles_AbsolutePath(t *testing.T) {
	r := validate.ValidateNoForbiddenFiles([]string{"/absolute/path"})
	if len(r.Findings) == 0 {
		t.Error("expected findings for absolute path")
	}
}
