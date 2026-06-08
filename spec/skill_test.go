package spec_test

import (
	"testing"

	"github.com/domehahn/sklib/spec"
)

func TestValidateSkillName(t *testing.T) {
	valid := []string{"my-skill", "abc", "foo-bar-baz", "a1b2", "skill123"}
	for _, v := range valid {
		if err := spec.ValidateSkillName(v); err != nil {
			t.Errorf("ValidateSkillName(%q) unexpected error: %v", v, err)
		}
	}

	invalid := []string{"", "My-Skill", "-skill", "skill-", "skill--name", "skill name", "SKILL", "foo_bar"}
	for _, v := range invalid {
		if err := spec.ValidateSkillName(v); err == nil {
			t.Errorf("ValidateSkillName(%q) expected error, got nil", v)
		}
	}
}

func TestValidateNamespace(t *testing.T) {
	valid := []string{"", "default", "my.org", "my-org", "my_org", "org.team.unit"}
	for _, v := range valid {
		if err := spec.ValidateNamespace(v); err != nil {
			t.Errorf("ValidateNamespace(%q) unexpected error: %v", v, err)
		}
	}

	invalid := []string{"-start", ".start", "end-", "end.", "UPPER"}
	for _, v := range invalid {
		if err := spec.ValidateNamespace(v); err == nil {
			t.Errorf("ValidateNamespace(%q) expected error, got nil", v)
		}
	}
}

func TestDefaultNamespace(t *testing.T) {
	if spec.DefaultNamespace("") != "default" {
		t.Error("DefaultNamespace(\"\") should return \"default\"")
	}
	if spec.DefaultNamespace("myorg") != "myorg" {
		t.Error("DefaultNamespace(\"myorg\") should return \"myorg\"")
	}
}

func TestDefaultEntrypoint(t *testing.T) {
	if spec.DefaultEntrypoint("") != "SKILL.md" {
		t.Error("DefaultEntrypoint(\"\") should return \"SKILL.md\"")
	}
	if spec.DefaultEntrypoint("README.md") != "README.md" {
		t.Error("DefaultEntrypoint(\"README.md\") should return \"README.md\"")
	}
}

func TestSkillSupportsPlatform(t *testing.T) {
	s := spec.Skill{
		Name:           "my-skill",
		CompatibleWith: []spec.Platform{spec.PlatformCodex, spec.PlatformCursor},
	}
	if !s.SupportsPlatform(spec.PlatformCodex) {
		t.Error("expected skill to support codex")
	}
	if s.SupportsPlatform(spec.PlatformOllama) {
		t.Error("expected skill to not support ollama")
	}
}

func TestSkillSupportsPlatform_All(t *testing.T) {
	s := spec.Skill{
		Name:           "universal-skill",
		CompatibleWith: []spec.Platform{spec.PlatformAll},
	}
	if !s.SupportsPlatform(spec.PlatformOllama) {
		t.Error("skill with 'all' should support any platform")
	}
}

func TestSkillRef(t *testing.T) {
	s := spec.Skill{
		Name:      "my-skill",
		Namespace: "myorg",
		Version:   "1.2.3",
	}
	ref := s.Ref()
	if ref.Namespace != "myorg" || ref.Name != "my-skill" || ref.Constraint != "1.2.3" {
		t.Errorf("unexpected Ref: %+v", ref)
	}
}

func TestSkillRef_DefaultNamespace(t *testing.T) {
	s := spec.Skill{Name: "my-skill", Version: "1.0.0"}
	ref := s.Ref()
	if ref.Namespace != "default" {
		t.Errorf("expected namespace 'default', got %q", ref.Namespace)
	}
}
