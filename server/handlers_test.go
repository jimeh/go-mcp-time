package server

import (
	"testing"
	"time"

	"github.com/jimeh/go-mcp-time/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCurrentTime(t *testing.T) {
	tests := []struct {
		name    string
		params  types.GetCurrentTimeParams
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid UTC timezone",
			params:  types.GetCurrentTimeParams{Timezone: "UTC"},
			wantErr: false,
		},
		{
			name:    "valid America/New_York timezone",
			params:  types.GetCurrentTimeParams{Timezone: "America/New_York"},
			wantErr: false,
		},
		{
			name:    "empty timezone defaults to UTC",
			params:  types.GetCurrentTimeParams{Timezone: ""},
			wantErr: false,
		},
		{
			name:    "invalid timezone",
			params:  types.GetCurrentTimeParams{Timezone: "Invalid/Timezone"},
			wantErr: true,
			errMsg:  "invalid timezone: Invalid/Timezone",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetCurrentTime(tt.params)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				assert.Nil(t, result)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)

			expectedTimezone := tt.params.Timezone
			if expectedTimezone == "" {
				expectedTimezone = "UTC"
			}
			assert.Equal(t, expectedTimezone, result.Timezone)
			assert.NotEmpty(t, result.Datetime)
			assert.Regexp(t, `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}$`, result.Datetime)
		})
	}
}

func TestConvertTime(t *testing.T) {
	tests := []struct {
		name    string
		params  types.ConvertTimeParams
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid time conversion UTC to EST",
			params: types.ConvertTimeParams{
				SourceTimezone: "UTC",
				Time:           "15:30",
				TargetTimezone: "America/New_York",
			},
			wantErr: false,
		},
		{
			name: "valid time conversion between timezones",
			params: types.ConvertTimeParams{
				SourceTimezone: "America/Los_Angeles",
				Time:           "09:00",
				TargetTimezone: "Europe/London",
			},
			wantErr: false,
		},
		{
			name: "invalid time format - missing colon",
			params: types.ConvertTimeParams{
				SourceTimezone: "UTC",
				Time:           "1530",
				TargetTimezone: "America/New_York",
			},
			wantErr: true,
			errMsg:  "invalid time format: 1530 (expected HH:MM)",
		},
		{
			name: "invalid time format - invalid hour",
			params: types.ConvertTimeParams{
				SourceTimezone: "UTC",
				Time:           "25:30",
				TargetTimezone: "America/New_York",
			},
			wantErr: true,
			errMsg:  "invalid time format: 25:30 (expected HH:MM)",
		},
		{
			name: "invalid time format - invalid minute",
			params: types.ConvertTimeParams{
				SourceTimezone: "UTC",
				Time:           "15:70",
				TargetTimezone: "America/New_York",
			},
			wantErr: true,
			errMsg:  "invalid time format: 15:70 (expected HH:MM)",
		},
		{
			name: "invalid source timezone",
			params: types.ConvertTimeParams{
				SourceTimezone: "Invalid/Source",
				Time:           "15:30",
				TargetTimezone: "America/New_York",
			},
			wantErr: true,
			errMsg:  "invalid source timezone: Invalid/Source",
		},
		{
			name: "invalid target timezone",
			params: types.ConvertTimeParams{
				SourceTimezone: "UTC",
				Time:           "15:30",
				TargetTimezone: "Invalid/Target",
			},
			wantErr: true,
			errMsg:  "invalid target timezone: Invalid/Target",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertTime(tt.params)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				assert.Nil(t, result)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)

			assert.Equal(t, tt.params.SourceTimezone, result.Source.Timezone)
			assert.Equal(t, tt.params.TargetTimezone, result.Target.Timezone)
			assert.NotEmpty(t, result.Source.Datetime)
			assert.NotEmpty(t, result.Target.Datetime)
			assert.NotEmpty(t, result.Source.Offset)
			assert.NotEmpty(t, result.Target.Offset)
			assert.NotEmpty(t, result.TimeDifference)

			assert.Regexp(t, `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}$`, result.Source.Datetime)
			assert.Regexp(t, `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[+-]\d{2}:\d{2}$`, result.Target.Datetime)
			assert.Regexp(t, `^[+-]\d{2}:\d{2}$`, result.Source.Offset)
			assert.Regexp(t, `^[+-]\d{2}:\d{2}$`, result.Target.Offset)
		})
	}
}

func TestFormatOffset(t *testing.T) {
	tests := []struct {
		name     string
		timezone string
		expected string
	}{
		{
			name:     "UTC timezone",
			timezone: "UTC",
			expected: "+00:00",
		},
		{
			name:     "EST timezone (negative offset)",
			timezone: "America/New_York",
			expected: "-05:00", // EST or -04:00 for EDT depending on time of year
		},
		{
			name:     "JST timezone (positive offset)",
			timezone: "Asia/Tokyo",
			expected: "+09:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loc, err := time.LoadLocation(tt.timezone)
			require.NoError(t, err)

			testTime := time.Date(2024, 1, 15, 12, 0, 0, 0, loc)
			result := formatOffset(testTime)

			if tt.timezone == "America/New_York" {
				assert.True(t, result == "-05:00" || result == "-04:00", "Expected EST (-05:00) or EDT (-04:00), got %s", result)
			} else {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestFormatTimeDifference(t *testing.T) {
	tests := []struct {
		name        string
		diffSeconds int
		expected    string
	}{
		{
			name:        "zero difference",
			diffSeconds: 0,
			expected:    "0 hours",
		},
		{
			name:        "1 hour positive",
			diffSeconds: 3600,
			expected:    "1 hour",
		},
		{
			name:        "1 hour negative",
			diffSeconds: -3600,
			expected:    "-1 hour",
		},
		{
			name:        "multiple hours positive",
			diffSeconds: 7200, // 2 hours
			expected:    "2 hours",
		},
		{
			name:        "multiple hours negative",
			diffSeconds: -7200, // -2 hours
			expected:    "-2 hours",
		},
		{
			name:        "1 minute positive",
			diffSeconds: 60,
			expected:    "1 minute",
		},
		{
			name:        "1 minute negative",
			diffSeconds: -60,
			expected:    "-1 minute",
		},
		{
			name:        "multiple minutes positive",
			diffSeconds: 300, // 5 minutes
			expected:    "5 minutes",
		},
		{
			name:        "multiple minutes negative",
			diffSeconds: -300, // -5 minutes
			expected:    "-5 minutes",
		},
		{
			name:        "hours and minutes positive",
			diffSeconds: 3900, // 1 hour 5 minutes
			expected:    "1 hour 5 minutes",
		},
		{
			name:        "hours and minutes negative",
			diffSeconds: -3900, // -1 hour -5 minutes
			expected:    "-1 hour -5 minutes",
		},
		{
			name:        "multiple hours and 1 minute",
			diffSeconds: 7260, // 2 hours 1 minute
			expected:    "2 hours 1 minute",
		},
		{
			name:        "1 hour and multiple minutes",
			diffSeconds: 3720, // 1 hour 2 minutes
			expected:    "1 hour 2 minutes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatTimeDifference(tt.diffSeconds)
			assert.Equal(t, tt.expected, result)
		})
	}
}
