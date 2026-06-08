package ref_test

import (
	"testing"

	"github.com/domehahn/sklib/ref"
	"github.com/domehahn/sklib/spec"
)

func TestParseSkillRef_Name(t *testing.T) {
	r, err := ref.ParseSkillRef("my-skill")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Namespace != "default" || r.Name != "my-skill" || r.Constraint != "" {
		t.Errorf("unexpected ref: %+v", r)
	}
}

func TestParseSkillRef_NameVersion(t *testing.T) {
	r, err := ref.ParseSkillRef("my-skill@1.2.3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Name != "my-skill" || r.Constraint != "1.2.3" {
		t.Errorf("unexpected ref: %+v", r)
	}
}

func TestParseSkillRef_NamespaceSlashName(t *testing.T) {
	r, err := ref.ParseSkillRef("myorg/my-skill")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Namespace != "myorg" || r.Name != "my-skill" {
		t.Errorf("unexpected ref: %+v", r)
	}
}

func TestParseSkillRef_Full(t *testing.T) {
	r, err := ref.ParseSkillRef("myorg/my-skill@^1.2.0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Namespace != "myorg" || r.Name != "my-skill" || r.Constraint != "^1.2.0" {
		t.Errorf("unexpected ref: %+v", r)
	}
}

func TestParseSkillRef_Latest(t *testing.T) {
	r, err := ref.ParseSkillRef("myorg/my-skill@latest")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Constraint != "latest" {
		t.Errorf("expected constraint 'latest', got %q", r.Constraint)
	}
}

func TestParseSkillRef_Invalid(t *testing.T) {
	cases := []string{
		"",
		"My-Skill",
		"/name",
		"ns/",
		"ns//name",
		"name@",
	}
	for _, tc := range cases {
		_, err := ref.ParseSkillRef(tc)
		if err == nil {
			t.Errorf("ParseSkillRef(%q) expected error, got nil", tc)
		}
	}
}

func TestParseSkillVersionRef_Valid(t *testing.T) {
	r, err := ref.ParseSkillVersionRef("myorg/my-skill@1.2.3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Namespace != "myorg" || r.Name != "my-skill" || r.Version != "1.2.3" {
		t.Errorf("unexpected ref: %+v", r)
	}
}

func TestParseSkillVersionRef_StripV(t *testing.T) {
	r, err := ref.ParseSkillVersionRef("my-skill@v2.0.1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Version != "2.0.1" {
		t.Errorf("expected version '2.0.1', got %q", r.Version)
	}
}

func TestParseSkillVersionRef_RejectRange(t *testing.T) {
	_, err := ref.ParseSkillVersionRef("my-skill@^1.2.0")
	if err == nil {
		t.Error("expected error for range constraint in version ref")
	}
}

func TestParseSkillVersionRef_RejectMissing(t *testing.T) {
	_, err := ref.ParseSkillVersionRef("my-skill")
	if err == nil {
		t.Error("expected error when version is missing")
	}
}

func TestParseSkillVersionRef_RejectLatest(t *testing.T) {
	_, err := ref.ParseSkillVersionRef("my-skill@latest")
	if err == nil {
		t.Error("expected error for 'latest' in version ref")
	}
}

func TestString(t *testing.T) {
	r := spec.SkillRef{Namespace: "myorg", Name: "my-skill", Constraint: "^1.0.0"}
	got := ref.String(r)
	if got != "myorg/my-skill@^1.0.0" {
		t.Errorf("unexpected String: %q", got)
	}
}

func TestString_NoConstraint(t *testing.T) {
	r := spec.SkillRef{Namespace: "myorg", Name: "my-skill"}
	got := ref.String(r)
	if got != "myorg/my-skill" {
		t.Errorf("unexpected String: %q", got)
	}
}

func TestVersionString(t *testing.T) {
	r := spec.SkillVersionRef{Namespace: "default", Name: "my-skill", Version: "1.0.0"}
	got := ref.VersionString(r)
	if got != "default/my-skill@1.0.0" {
		t.Errorf("unexpected VersionString: %q", got)
	}
}
