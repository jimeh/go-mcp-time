// Package server implements the MCP time server functionality.
package server

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jimeh/go-time-mcp/types"
)

var timeRegex = regexp.MustCompile(`^([01]\d|2[0-3]):([0-5]\d)$`)

// GetCurrentTime retrieves the current time for the specified timezone.
func GetCurrentTime(
	params types.GetCurrentTimeParams,
) (*types.TimeResult, error) {
	timezone := params.Timezone
	if timezone == "" {
		timezone = "UTC"
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, fmt.Errorf("invalid timezone: %s", timezone)
	}

	now := time.Now().In(loc)
	isDST := now.IsDST()

	result := &types.TimeResult{
		Timezone: timezone,
		Datetime: now.Format("2006-01-02T15:04:05-07:00"),
		IsDST:    isDST,
	}

	return result, nil
}

// ConvertTime converts a time from one timezone to another.
func ConvertTime(
	params types.ConvertTimeParams,
) (*types.TimeConversionResult, error) {
	if !timeRegex.MatchString(params.Time) {
		return nil, fmt.Errorf(
			"invalid time format: %s (expected HH:MM)", params.Time)
	}

	sourceLoc, err := time.LoadLocation(params.SourceTimezone)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid source timezone: %s", params.SourceTimezone)
	}

	targetLoc, err := time.LoadLocation(params.TargetTimezone)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid target timezone: %s", params.TargetTimezone)
	}

	timeParts := strings.Split(params.Time, ":")
	hour, _ := strconv.Atoi(timeParts[0])
	minute, _ := strconv.Atoi(timeParts[1])

	now := time.Now()
	sourceTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, sourceLoc)
	targetTime := sourceTime.In(targetLoc)

	sourceInfo := types.TimezoneInfo{
		Timezone: params.SourceTimezone,
		Datetime: sourceTime.Format("2006-01-02T15:04:05-07:00"),
		IsDST:    sourceTime.IsDST(),
		Offset:   formatOffset(sourceTime),
	}

	targetInfo := types.TimezoneInfo{
		Timezone: params.TargetTimezone,
		Datetime: targetTime.Format("2006-01-02T15:04:05-07:00"),
		IsDST:    targetTime.IsDST(),
		Offset:   formatOffset(targetTime),
	}

	_, sourceOffset := sourceTime.Zone()
	_, targetOffset := targetTime.Zone()
	diffSeconds := targetOffset - sourceOffset
	timeDifference := formatTimeDifference(diffSeconds)

	result := &types.TimeConversionResult{
		Source:         sourceInfo,
		Target:         targetInfo,
		TimeDifference: timeDifference,
	}

	return result, nil
}

func formatOffset(t time.Time) string {
	_, offset := t.Zone()
	hours := offset / 3600
	minutes := (offset % 3600) / 60

	if offset >= 0 {
		return fmt.Sprintf("+%02d:%02d", hours, minutes)
	}

	return fmt.Sprintf("-%02d:%02d", -hours, -minutes)
}

func formatTimeDifference(diffSeconds int) string {
	hours := diffSeconds / 3600
	minutes := (diffSeconds % 3600) / 60

	if diffSeconds == 0 {
		return "0 hours"
	}

	var parts []string
	if hours != 0 {
		if hours == 1 || hours == -1 {
			parts = append(parts, fmt.Sprintf("%d hour", hours))
		} else {
			parts = append(parts, fmt.Sprintf("%d hours", hours))
		}
	}

	if minutes != 0 {
		if minutes == 1 || minutes == -1 {
			parts = append(parts, fmt.Sprintf("%d minute", minutes))
		} else {
			parts = append(parts, fmt.Sprintf("%d minutes", minutes))
		}
	}

	if len(parts) == 0 {
		return "0 hours"
	}

	return strings.Join(parts, " ")
}
