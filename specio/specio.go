// Package specio provides helpers for reading and writing sklib spec files (YAML and JSON).
package specio

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/domehahn/sklib/spec"
)

// ReadSkillYAML reads and parses a skill.yaml file from path.
func ReadSkillYAML(path string) (*spec.Skill, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var s spec.Skill
	if err := yaml.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return &s, nil
}

// WriteSkillYAML serialises a Skill to a skill.yaml file at path.
func WriteSkillYAML(path string, s spec.Skill) error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Errorf("serialising skill yaml: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}
	return nil
}

// ReadAgentSkillsManifest reads and parses an agent-skills.yaml file from path.
func ReadAgentSkillsManifest(path string) (*spec.AgentSkillsManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var m spec.AgentSkillsManifest
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return &m, nil
}

// WriteAgentSkillsManifest serialises an AgentSkillsManifest to a YAML file at path.
func WriteAgentSkillsManifest(path string, m spec.AgentSkillsManifest) error {
	data, err := yaml.Marshal(m)
	if err != nil {
		return fmt.Errorf("serialising manifest yaml: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}
	return nil
}

// ReadAgentSkillsLock reads and parses an agent-skills.lock.yaml file from path.
func ReadAgentSkillsLock(path string) (*spec.AgentSkillsLock, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var l spec.AgentSkillsLock
	if err := yaml.Unmarshal(data, &l); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return &l, nil
}

// WriteAgentSkillsLock serialises an AgentSkillsLock to a YAML file at path.
func WriteAgentSkillsLock(path string, l spec.AgentSkillsLock) error {
	data, err := yaml.Marshal(l)
	if err != nil {
		return fmt.Errorf("serialising lockfile yaml: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}
	return nil
}

// ReadPackageManifest reads and parses a manifest.json file from path.
func ReadPackageManifest(path string) (*spec.PackageManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var m spec.PackageManifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return &m, nil
}

// WritePackageManifest serialises a PackageManifest to a JSON file at path.
func WritePackageManifest(path string, m spec.PackageManifest) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("serialising manifest json: %w", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", path, err)
	}
	return nil
}
