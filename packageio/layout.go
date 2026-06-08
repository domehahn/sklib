// Package packageio provides constants and helpers for skill package artifact layout.
package packageio

import "strings"

// Well-known filenames inside a skill package.
const (
	FileSkillMD   = "SKILL.md"
	FileVERSION   = "VERSION"
	FileSkillYAML = "skill.yaml"
	FileChangelog = "CHANGELOG.md"
	FileReadme    = "README.md"
	FileLicense   = "LICENSE"
	FileManifest  = "manifest.json"
	FileChecksums = "checksums.txt"
)

// requiredPackageFiles is the minimal set of files every valid package must contain.
var requiredPackageFiles = []string{
	FileSkillMD,
	FileSkillYAML,
	FileManifest,
	FileChecksums,
}

// forbiddenPathPrefixes lists path prefixes that must never appear in a package.
var forbiddenPathPrefixes = []string{
	".git/",
	".git",
	"../",
	".env",
}

// forbiddenPathSegments lists directory/file names that must never appear as path segments.
var forbiddenPathSegments = []string{
	".git",
	"node_modules",
	".venv",
	"__pycache__",
}

// generatedFiles are files produced by the packaging tool itself (not authored by the skill).
var generatedFiles = []string{
	FileManifest,
	FileChecksums,
}

// ValidatePackageLayout checks that paths contains all required files and no forbidden paths.
func ValidatePackageLayout(paths []string) error {
	pathSet := make(map[string]bool, len(paths))
	for _, p := range paths {
		pathSet[p] = true
	}

	for _, req := range requiredPackageFiles {
		if !pathSet[req] {
			return errMissingFile(req)
		}
	}

	for _, p := range paths {
		if IsForbiddenPackagePath(p) {
			return errForbiddenPath(p)
		}
	}
	return nil
}

// IsForbiddenPackagePath returns true if the path must not appear in a skill package.
func IsForbiddenPackagePath(path string) bool {
	// Reject absolute paths.
	if strings.HasPrefix(path, "/") {
		return true
	}

	// Reject path traversal.
	if strings.Contains(path, "../") || path == ".." {
		return true
	}

	// Reject forbidden prefixes.
	lower := strings.ToLower(path)
	for _, prefix := range forbiddenPathPrefixes {
		if strings.HasPrefix(lower, prefix) {
			return true
		}
	}

	// Reject paths containing forbidden segments.
	segments := strings.Split(path, "/")
	for _, seg := range segments {
		segLower := strings.ToLower(seg)
		for _, bad := range forbiddenPathSegments {
			if segLower == bad {
				return true
			}
		}
		// Reject obvious private key file extensions.
		if strings.HasSuffix(segLower, ".pem") ||
			strings.HasSuffix(segLower, ".key") ||
			strings.HasSuffix(segLower, ".pfx") ||
			strings.HasSuffix(segLower, ".p12") {
			return true
		}
	}

	return false
}

// IsGeneratedPackageFile returns true if the file is produced by the packaging tool,
// not authored by the skill developer.
func IsGeneratedPackageFile(path string) bool {
	for _, g := range generatedFiles {
		if path == g {
			return true
		}
	}
	return false
}

type missingFileError struct{ name string }

func (e missingFileError) Error() string {
	return "package is missing required file: " + e.name
}

func errMissingFile(name string) error { return missingFileError{name} }

type forbiddenPathError struct{ path string }

func (e forbiddenPathError) Error() string {
	return "package contains forbidden path: " + e.path
}

func errForbiddenPath(path string) error { return forbiddenPathError{path} }
