// Package ref provides parsing and formatting for skill coordinate references.
package ref

import (
	"fmt"
	"strings"

	"github.com/domehahn/sklib/spec"
)

// ParseSkillRef parses a skill reference string into a SkillRef.
//
// Supported forms:
//
//	name
//	name@constraint
//	namespace/name
//	namespace/name@constraint
//
// The namespace defaults to "default" when omitted.
// The constraint may be a version (e.g. "1.2.3"), a range (e.g. "^1.2.0"), or "latest".
func ParseSkillRef(value string) (spec.SkillRef, error) {
	if value == "" {
		return spec.SkillRef{}, fmt.Errorf("skill ref must not be empty")
	}

	ns, name, constraint, err := splitRef(value)
	if err != nil {
		return spec.SkillRef{}, err
	}

	if err := spec.ValidateSkillName(name); err != nil {
		return spec.SkillRef{}, fmt.Errorf("invalid skill ref %q: %w", value, err)
	}
	if err := spec.ValidateNamespace(ns); err != nil {
		return spec.SkillRef{}, fmt.Errorf("invalid skill ref %q: %w", value, err)
	}

	return spec.SkillRef{
		Namespace:  spec.DefaultNamespace(ns),
		Name:       name,
		Constraint: constraint,
	}, nil
}

// ParseSkillVersionRef parses a skill reference that must contain an exact SemVer version.
//
// Supported forms:
//
//	name@1.2.3
//	namespace/name@1.2.3
func ParseSkillVersionRef(value string) (spec.SkillVersionRef, error) {
	if value == "" {
		return spec.SkillVersionRef{}, fmt.Errorf("skill version ref must not be empty")
	}

	ns, name, constraint, err := splitRef(value)
	if err != nil {
		return spec.SkillVersionRef{}, err
	}

	if constraint == "" {
		return spec.SkillVersionRef{}, fmt.Errorf("skill version ref %q: version is required", value)
	}

	version, err := spec.NormalizeVersion(constraint)
	if err != nil {
		return spec.SkillVersionRef{}, fmt.Errorf("skill version ref %q: %w", value, err)
	}
	if !spec.IsStableSemVer(version) {
		return spec.SkillVersionRef{}, fmt.Errorf("skill version ref %q: version must be stable SemVer (no pre-release or build metadata)", value)
	}

	if err := spec.ValidateSkillName(name); err != nil {
		return spec.SkillVersionRef{}, fmt.Errorf("invalid skill version ref %q: %w", value, err)
	}
	if err := spec.ValidateNamespace(ns); err != nil {
		return spec.SkillVersionRef{}, fmt.Errorf("invalid skill version ref %q: %w", value, err)
	}

	return spec.SkillVersionRef{
		Namespace: spec.DefaultNamespace(ns),
		Name:      name,
		Version:   version,
	}, nil
}

// String formats a SkillRef back into its canonical string form.
func String(r spec.SkillRef) string {
	ns := r.Namespace
	if ns == "" {
		ns = spec.DefaultNamespaceValue
	}
	base := ns + "/" + r.Name
	if r.Constraint != "" {
		return base + "@" + r.Constraint
	}
	return base
}

// VersionString formats a SkillVersionRef into its canonical string form.
func VersionString(r spec.SkillVersionRef) string {
	ns := r.Namespace
	if ns == "" {
		ns = spec.DefaultNamespaceValue
	}
	return ns + "/" + r.Name + "@" + r.Version
}

// splitRef splits a raw ref string into (namespace, name, constraint).
// namespace is empty if not provided.
func splitRef(value string) (ns, name, constraint string, err error) {
	// Split off the version constraint first.
	atIdx := strings.Index(value, "@")
	if atIdx >= 0 {
		constraint = value[atIdx+1:]
		value = value[:atIdx]
		if constraint == "" {
			err = fmt.Errorf("skill ref has trailing '@' but no version")
			return
		}
	}

	// Split namespace from name.
	slashIdx := strings.Index(value, "/")
	if slashIdx >= 0 {
		ns = value[:slashIdx]
		name = value[slashIdx+1:]
		if ns == "" {
			err = fmt.Errorf("skill ref has empty namespace before '/'")
			return
		}
		if name == "" {
			err = fmt.Errorf("skill ref has empty name after '/'")
			return
		}
	} else {
		name = value
	}

	if name == "" {
		err = fmt.Errorf("skill ref has empty name")
	}
	return
}
