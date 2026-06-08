package validate

import (
	"strings"
)

// obviousSecretPatterns are filename substrings that suggest private credentials.
var obviousSecretPatterns = []string{
	".env",
	"id_rsa",
	"id_ed25519",
	"id_ecdsa",
	"id_dsa",
	".pem",
	".key",
	".pfx",
	".p12",
	"credentials",
	"secrets",
	"password",
	"passwd",
	"token",
	"api_key",
	"apikey",
	"private_key",
	"privatekey",
}

// warnOnPaths lists directory/filename substrings that should produce a warning.
var warnOnPaths = []string{
	"node_modules",
	".venv",
	"venv",
	"__pycache__",
	"dist/",
	"build/",
	".DS_Store",
	"Thumbs.db",
}

// ValidateNoForbiddenFiles returns findings for paths that contain obvious credential files.
func ValidateNoForbiddenFiles(paths []string) Result {
	var r Result
	for _, p := range paths {
		lower := strings.ToLower(p)
		if strings.HasPrefix(lower, "/") || strings.Contains(lower, "../") {
			r.AddPathWarning(p, "path is absolute or contains traversal: "+p)
			continue
		}
		for _, pat := range obviousSecretPatterns {
			if strings.Contains(lower, pat) {
				r.AddErrorCode("", "forbidden_file", "path matches forbidden secret pattern ("+pat+"): "+p)
				break
			}
		}
	}
	r.Finalize()
	return r
}

// ValidateNoObviousSecrets scans paths for filenames that suggest credential files.
// It is equivalent to ValidateNoForbiddenFiles but framed as a warning-only scan.
func ValidateNoObviousSecrets(paths []string) Result {
	var r Result
	for _, p := range paths {
		lower := strings.ToLower(p)
		for _, pat := range obviousSecretPatterns {
			if strings.Contains(lower, pat) {
				r.AddWarning("", "path may contain secrets ("+pat+"): "+p)
				break
			}
		}
	}
	r.Finalize()
	return r
}

// ValidateNoAbsoluteLocalPaths returns an error finding for any absolute path in the list.
func ValidateNoAbsoluteLocalPaths(paths []string) Result {
	var r Result
	for _, p := range paths {
		if strings.HasPrefix(p, "/") || strings.HasPrefix(p, "\\") {
			r.AddError("", "absolute path not allowed: "+p)
		}
		if strings.Contains(p, "../") || p == ".." {
			r.AddError("", "path traversal not allowed: "+p)
		}
	}
	r.Finalize()
	return r
}

// ValidatePackageLayout wraps packageio.ValidatePackageLayout into a Result.
// It also warns on build artifact and heavy-dependency paths.
func ValidatePackageLayout(paths []string) Result {
	var r Result

	// Check for forbidden paths.
	noForbidden := ValidateNoForbiddenFiles(paths)
	r.Findings = append(r.Findings, noForbidden.Findings...)

	// Warn on known large/build directories.
	for _, p := range paths {
		lower := strings.ToLower(p)
		for _, warn := range warnOnPaths {
			if strings.Contains(lower, warn) {
				r.AddPathWarning(p, "path looks like a build artifact or dependency directory: "+p)
				break
			}
		}
	}

	r.Finalize()
	return r
}
