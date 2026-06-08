package spec_test

import (
	"testing"

	"github.com/domehahn/sklib/spec"
)

func TestValidateManifestBasic_Valid(t *testing.T) {
	m := spec.AgentSkillsManifest{
		Version: 1,
		Registries: map[string]spec.RegistryConfig{
			"main": {Type: "skillforge", URL: "https://registry.example.com"},
		},
		Skills: []spec.ManifestSkill{
			{Name: "skill-a"},
			{Namespace: "myorg", Name: "skill-b"},
		},
	}
	if err := spec.ValidateManifestBasic(m); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidateManifestBasic_UnknownRegistryType(t *testing.T) {
	m := spec.AgentSkillsManifest{
		Registries: map[string]spec.RegistryConfig{
			"bad": {Type: "nosuchregistry"},
		},
	}
	if err := spec.ValidateManifestBasic(m); err == nil {
		t.Error("expected error for unknown registry type")
	}
}

func TestValidateManifestBasic_EmptyRegistryType(t *testing.T) {
	m := spec.AgentSkillsManifest{
		Registries: map[string]spec.RegistryConfig{
			"bad": {Type: ""},
		},
	}
	if err := spec.ValidateManifestBasic(m); err == nil {
		t.Error("expected error for empty registry type")
	}
}

func TestValidateManifestBasic_DuplicateSkill(t *testing.T) {
	m := spec.AgentSkillsManifest{
		Skills: []spec.ManifestSkill{
			{Name: "skill-a"},
			{Name: "skill-a"},
		},
	}
	if err := spec.ValidateManifestBasic(m); err == nil {
		t.Error("expected error for duplicate skill")
	}
}

func TestValidateManifestBasic_EmptySkillName(t *testing.T) {
	m := spec.AgentSkillsManifest{
		Skills: []spec.ManifestSkill{{Name: ""}},
	}
	if err := spec.ValidateManifestBasic(m); err == nil {
		t.Error("expected error for empty skill name")
	}
}

func TestNormalizeManifestDefaults(t *testing.T) {
	m := spec.AgentSkillsManifest{
		Skills: []spec.ManifestSkill{{Name: "skill-a"}},
	}
	spec.NormalizeManifestDefaults(&m)
	if m.Version != 1 {
		t.Errorf("expected version 1, got %d", m.Version)
	}
	if m.Skills[0].Namespace != "default" {
		t.Errorf("expected namespace 'default', got %q", m.Skills[0].Namespace)
	}
}

func TestIsKnownRegistryType(t *testing.T) {
	valid := []string{"skillforge", "artifactory", "gitlab", "github", "local", "generic-http"}
	for _, v := range valid {
		if !spec.IsKnownRegistryType(v) {
			t.Errorf("IsKnownRegistryType(%q) = false, want true", v)
		}
	}
	if spec.IsKnownRegistryType("unknown") {
		t.Error("IsKnownRegistryType(\"unknown\") = true, want false")
	}
}

func TestValidateManifestBasic_KnownRegistryTypes(t *testing.T) {
	for _, rt := range spec.KnownRegistryTypes {
		m := spec.AgentSkillsManifest{
			Registries: map[string]spec.RegistryConfig{
				"r": {Type: rt},
			},
		}
		if err := spec.ValidateManifestBasic(m); err != nil {
			t.Errorf("ValidateManifestBasic with type %q: unexpected error: %v", rt, err)
		}
	}
}
