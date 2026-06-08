package packageio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/domehahn/sklib/spec"
)

// ReadManifest reads and parses a manifest.json file from the given path.
func ReadManifest(path string) (*spec.PackageManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading manifest: %w", err)
	}
	var m spec.PackageManifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parsing manifest: %w", err)
	}
	return &m, nil
}

// WriteManifest serialises a PackageManifest to a manifest.json file.
func WriteManifest(path string, m spec.PackageManifest) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("serialising manifest: %w", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("writing manifest: %w", err)
	}
	return nil
}
