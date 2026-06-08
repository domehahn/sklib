package registryapi

// ErrorResponse is the standard error body returned by the registry API.
type ErrorResponse struct {
	Error   string            `json:"error"`
	Code    string            `json:"code,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

// Well-known registry API error codes.
const (
	CodeNotFound       = "not_found"
	CodeAlreadyExists  = "already_exists"
	CodeInvalidRequest = "invalid_request"
	CodeUnauthorized   = "unauthorized"
	CodeForbidden      = "forbidden"
	CodeYanked         = "yanked"
	CodeDeprecated     = "deprecated"
	CodeChecksumFailed = "checksum_failed"
	CodeInternal       = "internal_error"
)

// NewErrorResponse constructs an ErrorResponse with a code and human-readable message.
func NewErrorResponse(code, message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
		Code:  code,
	}
}
