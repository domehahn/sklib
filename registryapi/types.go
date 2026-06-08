// Package registryapi defines DTOs for the vendor-neutral Skill Registry API.
// It contains only data types — no HTTP clients.
package registryapi

import "github.com/domehahn/sklib/spec"

// CapabilitiesResponse is returned by GET /capabilities.
type CapabilitiesResponse struct {
	APIVersion          string               `json:"api_version"`
	Features            RegistryCapabilities `json:"features"`
	PackageTypes        []string             `json:"package_types,omitempty"`
	MaxPackageSizeBytes int64                `json:"max_package_size_bytes,omitempty"`
}

// RegistryCapabilities describes which operations a registry supports.
type RegistryCapabilities struct {
	Resolve           bool `json:"resolve"`
	Download          bool `json:"download"`
	Search            bool `json:"search"`
	Info              bool `json:"info"`
	ListVersions      bool `json:"list_versions"`
	Publish           bool `json:"publish"`
	Deprecate         bool `json:"deprecate"`
	Yank              bool `json:"yank"`
	Unyank            bool `json:"unyank"`
	SemVerConstraints bool `json:"semver_constraints"`
	Checksums         bool `json:"checksums"`
	Signatures        bool `json:"signatures"`
	Provenance        bool `json:"provenance"`
}

// SkillInfoResponse is returned by GET /skills/{namespace}/{name}.
type SkillInfoResponse struct {
	Namespace     string         `json:"namespace"`
	Name          string         `json:"name"`
	LatestVersion string         `json:"latest_version,omitempty"`
	Description   string         `json:"description,omitempty"`
	Tags          []string       `json:"tags,omitempty"`
	Owners        []string       `json:"owners,omitempty"`
	Versions      []SkillVersion `json:"versions,omitempty"`
}

// SkillVersion describes a single published version of a skill.
type SkillVersion struct {
	Version           string         `json:"version"`
	Deprecated        bool           `json:"deprecated,omitempty"`
	DeprecationReason string         `json:"deprecation_reason,omitempty"`
	Yanked            bool           `json:"yanked,omitempty"`
	YankReason        string         `json:"yank_reason,omitempty"`
	SHA256            string         `json:"sha256,omitempty"`
	PackageType       string         `json:"package_type,omitempty"`
	Artifact          string         `json:"artifact,omitempty"`
	DownloadURL       string         `json:"download_url,omitempty"`
	CompatibleWith    []spec.Platform `json:"compatible_with,omitempty"`
	CreatedAt         string         `json:"created_at,omitempty"`
}

// SearchResponse is returned by GET /skills?q=...
type SearchResponse struct {
	Results []SearchResult `json:"results"`
	Total   int            `json:"total,omitempty"`
}

// SearchResult is one item in a search response.
type SearchResult struct {
	Namespace     string   `json:"namespace"`
	Name          string   `json:"name"`
	LatestVersion string   `json:"latest_version,omitempty"`
	Description   string   `json:"description,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

// ResolveRequest represents query parameters for GET /skills/{namespace}/{name}/resolve.
type ResolveRequest struct {
	Constraint string `json:"constraint,omitempty"`
}

// ResolveResponse is returned by GET /skills/{namespace}/{name}/resolve.
type ResolveResponse struct {
	Namespace      string          `json:"namespace"`
	Name           string          `json:"name"`
	Version        string          `json:"version"`
	DownloadURL    string          `json:"download_url"`
	SHA256         string          `json:"sha256,omitempty"`
	PackageType    string          `json:"package_type,omitempty"`
	Artifact       string          `json:"artifact,omitempty"`
	CompatibleWith []spec.Platform `json:"compatible_with,omitempty"`
}

// PublishResponse is returned by PUT /skills/{namespace}/{name}/versions/{version}.
type PublishResponse struct {
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	DownloadURL string `json:"download_url,omitempty"`
	SHA256      string `json:"sha256,omitempty"`
	Created     bool   `json:"created"`
}

// GovernanceRequest is used for deprecate, yank, and unyank operations.
type GovernanceRequest struct {
	Reason string `json:"reason,omitempty"`
}
