package validate

import (
	"github.com/domehahn/sklib/packageio"
	"github.com/domehahn/sklib/spec"
)

// ValidateSkillFiles checks that all required package files are present and no forbidden paths exist.
func ValidateSkillFiles(paths []string, opts Options) Result {
	var r Result

	if err := packageio.ValidatePackageLayout(paths); err != nil {
		r.AddError("layout", err.Error())
	}

	fsResult := ValidatePackageLayout(paths)
	r.Findings = append(r.Findings, fsResult.Findings...)

	r.Finalize()
	return r
}

// ValidatePackageManifest validates a PackageManifest for required fields.
func ValidatePackageManifest(m spec.PackageManifest, opts Options) Result {
	var r Result

	if err := spec.ValidateSkillName(m.Name); err != nil {
		r.AddError("name", err.Error())
	}
	if m.Namespace != "" {
		if err := spec.ValidateNamespace(m.Namespace); err != nil {
			r.AddError("namespace", err.Error())
		}
	}
	if m.Version == "" {
		r.AddError("version", "version must not be empty")
	} else if !spec.IsStableSemVer(m.Version) {
		r.AddError("version", "package version must be stable SemVer: "+m.Version)
	}
	if m.Entrypoint == "" {
		r.AddError("entrypoint", "entrypoint must not be empty")
	}
	if m.PackageType == "" {
		r.AddError("package_type", "package_type must not be empty")
	}
	if m.SHA256 != "" {
		if err := spec.ValidateSHA256(m.SHA256); err != nil {
			r.AddError("sha256", err.Error())
		}
	}

	r.Finalize()
	return r
}
