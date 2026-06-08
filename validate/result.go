// Package validate provides composable validation primitives and result types for sklib.
package validate

// Severity indicates the importance of a validation finding.
type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
	SeverityInfo    Severity = "info"
)

// Finding is a single validation observation.
type Finding struct {
	Severity Severity `json:"severity"`
	Field    string   `json:"field,omitempty"`
	Path     string   `json:"path,omitempty"`
	Message  string   `json:"message"`
	Code     string   `json:"code,omitempty"`
}

// Result accumulates findings from one or more validators.
type Result struct {
	Valid    bool      `json:"valid"`
	Findings []Finding `json:"findings,omitempty"`
}

// AddError appends an error-severity finding.
func (r *Result) AddError(field, message string) {
	r.Findings = append(r.Findings, Finding{
		Severity: SeverityError,
		Field:    field,
		Message:  message,
	})
}

// AddErrorCode appends an error-severity finding with a machine-readable code.
func (r *Result) AddErrorCode(field, code, message string) {
	r.Findings = append(r.Findings, Finding{
		Severity: SeverityError,
		Field:    field,
		Code:     code,
		Message:  message,
	})
}

// AddWarning appends a warning-severity finding.
func (r *Result) AddWarning(field, message string) {
	r.Findings = append(r.Findings, Finding{
		Severity: SeverityWarning,
		Field:    field,
		Message:  message,
	})
}

// AddPathWarning appends a warning-severity finding referencing a file path.
func (r *Result) AddPathWarning(path, message string) {
	r.Findings = append(r.Findings, Finding{
		Severity: SeverityWarning,
		Path:     path,
		Message:  message,
	})
}

// AddInfo appends an info-severity finding.
func (r *Result) AddInfo(field, message string) {
	r.Findings = append(r.Findings, Finding{
		Severity: SeverityInfo,
		Field:    field,
		Message:  message,
	})
}

// HasErrors returns true if any finding has error severity.
func (r *Result) HasErrors() bool {
	for _, f := range r.Findings {
		if f.Severity == SeverityError {
			return true
		}
	}
	return false
}

// Finalize sets Valid based on whether any errors are present.
// Call this after all findings have been added.
func (r *Result) Finalize() {
	r.Valid = !r.HasErrors()
}

// Options controls strictness of validation.
type Options struct {
	// Strict enables additional checks that are warnings in permissive mode but errors in strict mode.
	Strict bool
}
