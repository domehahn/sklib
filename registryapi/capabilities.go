package registryapi

// DefaultCapabilities returns a CapabilitiesResponse with only the most
// basic operations enabled. Registry implementations can build on this.
func DefaultCapabilities(apiVersion string) CapabilitiesResponse {
	return CapabilitiesResponse{
		APIVersion: apiVersion,
		Features: RegistryCapabilities{
			Resolve:  true,
			Download: true,
			Info:     true,
			Search:   false,
		},
	}
}

// FullCapabilities returns a CapabilitiesResponse with all features enabled.
func FullCapabilities(apiVersion string) CapabilitiesResponse {
	return CapabilitiesResponse{
		APIVersion: apiVersion,
		Features: RegistryCapabilities{
			Resolve:           true,
			Download:          true,
			Search:            true,
			Info:              true,
			ListVersions:      true,
			Publish:           true,
			Deprecate:         true,
			Yank:              true,
			Unyank:            true,
			SemVerConstraints: true,
			Checksums:         true,
			Signatures:        false,
			Provenance:        false,
		},
	}
}
