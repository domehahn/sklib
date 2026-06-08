package spec_test

import (
	"testing"

	"github.com/domehahn/sklib/spec"
)

func TestNormalizeVersion(t *testing.T) {
	cases := []struct {
		in      string
		want    string
		wantErr bool
	}{
		{"v1.2.3", "1.2.3", false},
		{"1.2.3", "1.2.3", false},
		{"0.0.1", "0.0.1", false},
		{"1.0.0-alpha", "1.0.0-alpha", false},
		{"1.0.0+build.123", "1.0.0+build.123", false},
		{"1.0.0-beta+exp.sha.5114f85", "1.0.0-beta+exp.sha.5114f85", false},
		{"1.2", "", true},
		{"latest", "", true},
		{"", "", true},
		{"v", "", true},
		{"1.2.3.4", "", true},
		{"01.2.3", "", true},
	}
	for _, tc := range cases {
		got, err := spec.NormalizeVersion(tc.in)
		if (err != nil) != tc.wantErr {
			t.Errorf("NormalizeVersion(%q) error = %v, wantErr = %v", tc.in, err, tc.wantErr)
			continue
		}
		if !tc.wantErr && got != tc.want {
			t.Errorf("NormalizeVersion(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestIsStableSemVer(t *testing.T) {
	stable := []string{"1.2.3", "v1.2.3", "0.0.0", "10.20.30"}
	for _, v := range stable {
		if !spec.IsStableSemVer(v) {
			t.Errorf("IsStableSemVer(%q) = false, want true", v)
		}
	}
	notStable := []string{"1.2.3-alpha", "1.2.3+build", "1.2", "latest", ""}
	for _, v := range notStable {
		if spec.IsStableSemVer(v) {
			t.Errorf("IsStableSemVer(%q) = true, want false", v)
		}
	}
}

func TestIsSemVer(t *testing.T) {
	valid := []string{"1.2.3", "v1.2.3", "1.0.0-alpha", "1.0.0+001", "1.0.0-beta+exp"}
	for _, v := range valid {
		if !spec.IsSemVer(v) {
			t.Errorf("IsSemVer(%q) = false, want true", v)
		}
	}
	invalid := []string{"1.2", "latest", "", "1.2.3.4"}
	for _, v := range invalid {
		if spec.IsSemVer(v) {
			t.Errorf("IsSemVer(%q) = true, want false", v)
		}
	}
}

func TestCompareVersions(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"1.0.0", "1.0.0", 0},
		{"1.0.0", "2.0.0", -1},
		{"2.0.0", "1.0.0", 1},
		{"1.1.0", "1.0.9", 1},
		{"1.0.1", "1.0.2", -1},
		{"10.0.0", "9.9.9", 1},
	}
	for _, tc := range cases {
		got := spec.CompareVersions(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("CompareVersions(%q, %q) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}
