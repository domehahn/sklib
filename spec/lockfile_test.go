package spec_test

import (
	"strings"
	"testing"

	"github.com/domehahn/sklib/spec"
)

const validSHA256 = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

func TestValidateSHA256_Valid(t *testing.T) {
	if err := spec.ValidateSHA256(validSHA256); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateSHA256_Empty(t *testing.T) {
	if err := spec.ValidateSHA256(""); err == nil {
		t.Error("expected error for empty SHA256")
	}
}

func TestValidateSHA256_Short(t *testing.T) {
	if err := spec.ValidateSHA256("abc123"); err == nil {
		t.Error("expected error for short SHA256")
	}
}

func TestValidateSHA256_Uppercase(t *testing.T) {
	upper := strings.ToUpper(validSHA256)
	if err := spec.ValidateSHA256(upper); err == nil {
		t.Error("expected error for uppercase SHA256")
	}
}

func TestValidateSHA256_NonHex(t *testing.T) {
	bad := strings.Repeat("g", 64)
	if err := spec.ValidateSHA256(bad); err == nil {
		t.Error("expected error for non-hex SHA256")
	}
}

func TestValidateLockBasic_Valid(t *testing.T) {
	lock := spec.AgentSkillsLock{
		Version: 1,
		Skills: []spec.LockedSkill{
			{
				Name:    "my-skill",
				Version: "1.0.0",
				Source:  "skillforge://registry.example.com/default/my-skill@1.0.0",
				SHA256:  validSHA256,
			},
		},
	}
	if err := spec.ValidateLockBasic(lock); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateLockBasic_AllowPrerelease(t *testing.T) {
	lock := spec.AgentSkillsLock{
		Skills: []spec.LockedSkill{
			{
				Name:    "my-skill",
				Version: "1.0.0-alpha",
				Source:  "skillforge://registry.example.com/default/my-skill",
				SHA256:  validSHA256,
			},
		},
	}
	// Schema allows pre-release versions in lockfiles; only ranges and 'latest' are forbidden.
	if err := spec.ValidateLockBasic(lock); err != nil {
		t.Errorf("pre-release version should be allowed in lockfile: %v", err)
	}
}

func TestValidateLockBasic_RejectVersionRange(t *testing.T) {
	lock := spec.AgentSkillsLock{
		Skills: []spec.LockedSkill{
			{
				Name:    "my-skill",
				Version: "latest",
				Source:  "skillforge://registry.example.com/default/my-skill",
				SHA256:  validSHA256,
			},
		},
	}
	if err := spec.ValidateLockBasic(lock); err == nil {
		t.Error("expected error for 'latest' in lockfile version")
	}
}

func TestValidateLockBasic_RejectMissingSource(t *testing.T) {
	lock := spec.AgentSkillsLock{
		Skills: []spec.LockedSkill{
			{Name: "my-skill", Version: "1.0.0", SHA256: validSHA256},
		},
	}
	if err := spec.ValidateLockBasic(lock); err == nil {
		t.Error("expected error for missing source")
	}
}

func TestValidateLockBasic_RejectMissingSHA(t *testing.T) {
	lock := spec.AgentSkillsLock{
		Skills: []spec.LockedSkill{
			{Name: "my-skill", Version: "1.0.0", Source: "skillforge://example.com/my-skill"},
		},
	}
	if err := spec.ValidateLockBasic(lock); err == nil {
		t.Error("expected error for missing sha256")
	}
}

func TestValidateLockBasic_RejectEmptyName(t *testing.T) {
	lock := spec.AgentSkillsLock{
		Skills: []spec.LockedSkill{
			{Name: "", Version: "1.0.0", Source: "skillforge://x", SHA256: validSHA256},
		},
	}
	if err := spec.ValidateLockBasic(lock); err == nil {
		t.Error("expected error for empty name")
	}
}
