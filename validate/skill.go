package validate

import (
	"github.com/domehahn/sklib/spec"
)

// ValidateSkillMetadata validates the metadata fields of a Skill struct.
// In strict mode, description and compatible_with are required.
func ValidateSkillMetadata(s spec.Skill, opts Options) Result {
	var r Result

	if err := spec.ValidateSkillName(s.Name); err != nil {
		r.AddError("name", err.Error())
	}
	if s.Namespace != "" {
		if err := spec.ValidateNamespace(s.Namespace); err != nil {
			r.AddError("namespace", err.Error())
		}
	}
	if s.Version == "" {
		r.AddError("version", "version must not be empty")
	} else if !spec.IsSemVer(s.Version) {
		r.AddError("version", "version must be valid SemVer: "+s.Version)
	}

	if opts.Strict {
		if s.Description == "" {
			r.AddError("description", "description must not be empty (strict mode)")
		}
		if len(s.CompatibleWith) == 0 {
			r.AddError("compatible_with", "compatible_with must not be empty (strict mode)")
		}
	} else {
		if s.Description == "" {
			r.AddWarning("description", "description is empty")
		}
		if len(s.CompatibleWith) == 0 {
			r.AddWarning("compatible_with", "compatible_with is empty")
		}
	}

	r.Finalize()
	return r
}

// ValidateManifestSkill validates a single ManifestSkill entry.
func ValidateManifestSkill(s spec.ManifestSkill, opts Options) Result {
	var r Result

	if s.Name == "" {
		r.AddError("name", "name must not be empty")
	} else if err := spec.ValidateSkillName(s.Name); err != nil {
		r.AddError("name", err.Error())
	}
	if s.Namespace != "" {
		if err := spec.ValidateNamespace(s.Namespace); err != nil {
			r.AddError("namespace", err.Error())
		}
	}
	if s.Version != "" && !spec.IsSemVer(s.Version) {
		r.AddWarning("version", "version is not valid SemVer: "+s.Version)
	}

	for i, p := range s.Platforms {
		if !spec.IsKnownPlatform(string(p)) {
			r.AddErrorCode("platforms", "unknown_platform",
				"platforms["+itoa(i)+"]: unknown platform "+string(p))
		}
	}

	r.Finalize()
	return r
}

// ValidateLockEntry validates a single LockedSkill entry.
func ValidateLockEntry(s spec.LockedSkill, opts Options) Result {
	var r Result

	if s.Name == "" {
		r.AddError("name", "name must not be empty")
	}
	if s.Version == "" {
		r.AddError("version", "version must not be empty")
	} else if !spec.IsStableSemVer(s.Version) {
		r.AddError("version", "locked version must be stable SemVer (no pre-release or build metadata): "+s.Version)
	}
	if s.Source == "" {
		r.AddError("source", "source must not be empty")
	}
	if err := spec.ValidateSHA256(s.SHA256); err != nil {
		r.AddError("sha256", err.Error())
	}

	r.Finalize()
	return r
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	buf := [20]byte{}
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[pos:])
}
