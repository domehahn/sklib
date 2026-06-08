package packageio_test

import (
	"testing"

	"github.com/domehahn/sklib/packageio"
)

func TestValidatePackageLayout_Valid(t *testing.T) {
	paths := []string{
		packageio.FileSkillMD,
		packageio.FileSkillYAML,
		packageio.FileManifest,
		packageio.FileChecksums,
	}
	if err := packageio.ValidatePackageLayout(paths); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestValidatePackageLayout_MissingRequired(t *testing.T) {
	paths := []string{packageio.FileSkillMD, packageio.FileSkillYAML, packageio.FileManifest}
	if err := packageio.ValidatePackageLayout(paths); err == nil {
		t.Error("expected error for missing checksums.txt")
	}
}

func TestIsForbiddenPackagePath(t *testing.T) {
	forbidden := []string{
		"/absolute/path",
		"../traversal",
		".git/config",
		".git",
		"sub/.git/HEAD",
		"node_modules/foo",
		"secrets/key.pem",
		"certs/server.key",
	}
	for _, p := range forbidden {
		if !packageio.IsForbiddenPackagePath(p) {
			t.Errorf("IsForbiddenPackagePath(%q) = false, want true", p)
		}
	}
}

func TestIsForbiddenPackagePath_Allowed(t *testing.T) {
	allowed := []string{
		"SKILL.md",
		"skill.yaml",
		"manifest.json",
		"docs/usage.md",
		"src/main.go",
	}
	for _, p := range allowed {
		if packageio.IsForbiddenPackagePath(p) {
			t.Errorf("IsForbiddenPackagePath(%q) = true, want false", p)
		}
	}
}

func TestIsGeneratedPackageFile(t *testing.T) {
	if !packageio.IsGeneratedPackageFile(packageio.FileManifest) {
		t.Error("manifest.json should be a generated file")
	}
	if !packageio.IsGeneratedPackageFile(packageio.FileChecksums) {
		t.Error("checksums.txt should be a generated file")
	}
	if packageio.IsGeneratedPackageFile(packageio.FileSkillMD) {
		t.Error("SKILL.md should not be a generated file")
	}
}
