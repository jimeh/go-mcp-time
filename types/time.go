// Package types defines data structures for MCP time server operations.
package types

// GetCurrentTimeParams holds parameters for getting current time in a timezone.
type GetCurrentTimeParams struct {
	Timezone string `json:"timezone"`
}

// TimeResult contains the result of a current time request.
type TimeResult struct {
	Timezone string `json:"timezone"`
	Datetime string `json:"datetime"`
	IsDST    bool   `json:"is_dst"`
}

// ConvertTimeParams holds parameters for converting time between timezones.
type ConvertTimeParams struct {
	SourceTimezone string `json:"source_timezone"`
	Time           string `json:"time"`
	TargetTimezone string `json:"target_timezone"`
}

// TimezoneInfo contains timezone-specific information for a time.
type TimezoneInfo struct {
	Timezone string `json:"timezone"`
	Datetime string `json:"datetime"`
	IsDST    bool   `json:"is_dst"`
	Offset   string `json:"offset"`
}

// TimeConversionResult contains the result of a time conversion operation.
type TimeConversionResult struct {
	Source         TimezoneInfo `json:"source"`
	Target         TimezoneInfo `json:"target"`
	TimeDifference string       `json:"time_difference"`
}
