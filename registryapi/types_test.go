package registryapi_test

import (
	"encoding/json"
	"testing"

	"github.com/domehahn/sklib/registryapi"
	"github.com/domehahn/sklib/spec"
)

func TestCapabilitiesResponse_MarshalUnmarshal(t *testing.T) {
	caps := registryapi.FullCapabilities("v1")
	data, err := json.Marshal(caps)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	var got registryapi.CapabilitiesResponse
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if got.APIVersion != "v1" {
		t.Errorf("expected api_version 'v1', got %q", got.APIVersion)
	}
	if !got.Features.Resolve {
		t.Error("expected Resolve to be true")
	}
}

func TestDefaultCapabilities(t *testing.T) {
	caps := registryapi.DefaultCapabilities("v1")
	if !caps.Features.Resolve || !caps.Features.Download {
		t.Error("default capabilities should include resolve and download")
	}
	if caps.Features.Publish {
		t.Error("default capabilities should not include publish")
	}
}

func TestResolveResponse_MarshalUnmarshal(t *testing.T) {
	resp := registryapi.ResolveResponse{
		Namespace:      "myorg",
		Name:           "my-skill",
		Version:        "1.2.3",
		DownloadURL:    "https://example.com/my-skill-1.2.3.tar.gz",
		CompatibleWith: []spec.Platform{spec.PlatformCodex},
	}
	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	var got registryapi.ResolveResponse
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if got.Version != "1.2.3" {
		t.Errorf("expected version 1.2.3, got %q", got.Version)
	}
	if len(got.CompatibleWith) != 1 || got.CompatibleWith[0] != spec.PlatformCodex {
		t.Errorf("unexpected CompatibleWith: %v", got.CompatibleWith)
	}
}

func TestErrorResponse(t *testing.T) {
	e := registryapi.NewErrorResponse(registryapi.CodeNotFound, "skill not found")
	if e.Code != registryapi.CodeNotFound {
		t.Errorf("unexpected code: %q", e.Code)
	}
	if e.Error == "" {
		t.Error("error message must not be empty")
	}
}

func TestPublishResponse_MarshalUnmarshal(t *testing.T) {
	resp := registryapi.PublishResponse{
		Namespace: "myorg",
		Name:      "my-skill",
		Version:   "2.0.0",
		Created:   true,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	var got registryapi.PublishResponse
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if !got.Created {
		t.Error("expected Created=true")
	}
}
