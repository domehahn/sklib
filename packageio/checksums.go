package packageio

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/domehahn/sklib/spec"
)

// ParseChecksumsText parses a checksums.txt file in the format "<sha256>  <path>" (two spaces).
func ParseChecksumsText(data []byte) ([]spec.ChecksumEntry, error) {
	var entries []spec.ChecksumEntry
	scanner := bufio.NewScanner(bytes.NewReader(data))
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Format: "<sha256>  <path>"
		idx := strings.Index(line, "  ")
		if idx < 0 {
			return nil, fmt.Errorf("checksums.txt line %d: expected \"<sha256>  <path>\", got %q", lineNo, line)
		}
		hash := line[:idx]
		path := line[idx+2:]
		if err := spec.ValidateSHA256(hash); err != nil {
			return nil, fmt.Errorf("checksums.txt line %d: %w", lineNo, err)
		}
		if path == "" {
			return nil, fmt.Errorf("checksums.txt line %d: empty file path", lineNo)
		}
		entries = append(entries, spec.ChecksumEntry{Path: path, SHA256: hash})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

// FormatChecksumsText serialises a slice of ChecksumEntry values to checksums.txt format.
func FormatChecksumsText(entries []spec.ChecksumEntry) []byte {
	var buf bytes.Buffer
	for _, e := range entries {
		buf.WriteString(e.SHA256)
		buf.WriteString("  ")
		buf.WriteString(e.Path)
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}
