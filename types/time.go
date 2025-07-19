package types

type GetCurrentTimeParams struct {
	Timezone string `json:"timezone"`
}

type TimeResult struct {
	Timezone string `json:"timezone"`
	Datetime string `json:"datetime"`
	IsDST    bool   `json:"is_dst"`
}

type ConvertTimeParams struct {
	SourceTimezone string `json:"source_timezone"`
	Time           string `json:"time"`
	TargetTimezone string `json:"target_timezone"`
}

type TimezoneInfo struct {
	Timezone string `json:"timezone"`
	Datetime string `json:"datetime"`
	IsDST    bool   `json:"is_dst"`
	Offset   string `json:"offset"`
}

type TimeConversionResult struct {
	Source         TimezoneInfo `json:"source"`
	Target         TimezoneInfo `json:"target"`
	TimeDifference string       `json:"time_difference"`
}