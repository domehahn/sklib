package packageio_test

import (
	"strings"
	"testing"

	"github.com/domehahn/sklib/packageio"
	"github.com/domehahn/sklib/spec"
)

const sha256a = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
const sha256b = "abc123def456789012345678901234567890123456789012345678901234abcd"

func TestFormatAndParseChecksums(t *testing.T) {
	entries := []spec.ChecksumEntry{
		{Path: "SKILL.md", SHA256: sha256a},
		{Path: "skill.yaml", SHA256: sha256b},
	}
	data := packageio.FormatChecksumsText(entries)
	got, err := packageio.ParseChecksumsText(data)
	if err != nil {
		t.Fatalf("ParseChecksumsText error: %v", err)
	}
	if len(got) != len(entries) {
		t.Fatalf("expected %d entries, got %d", len(entries), len(got))
	}
	for i, e := range entries {
		if got[i].Path != e.Path || got[i].SHA256 != e.SHA256 {
			t.Errorf("entry %d mismatch: got %+v, want %+v", i, got[i], e)
		}
	}
}

func TestParseChecksumsText_Comment(t *testing.T) {
	input := "# comment\n" + sha256a + "  SKILL.md\n"
	got, err := packageio.ParseChecksumsText([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || got[0].Path != "SKILL.md" {
		t.Errorf("unexpected result: %+v", got)
	}
}

func TestParseChecksumsText_BadFormat(t *testing.T) {
	input := "badsingleline\n"
	_, err := packageio.ParseChecksumsText([]byte(input))
	if err == nil {
		t.Error("expected error for bad format")
	}
}

func TestParseChecksumsText_BadHash(t *testing.T) {
	input := strings.Repeat("z", 64) + "  SKILL.md\n"
	_, err := packageio.ParseChecksumsText([]byte(input))
	if err == nil {
		t.Error("expected error for invalid hash")
	}
}
