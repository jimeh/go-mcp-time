package server

import (
	"context"
	"fmt"

	"github.com/jimeh/go-time-mcp/types"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type TimeServer struct {
	server *server.MCPServer
}

func NewTimeServer(localTimezone string) (*TimeServer, error) {
	s := server.NewMCPServer("time-server", "1.0.0")

	ts := &TimeServer{
		server: s,
	}

	err := ts.registerTools(localTimezone)
	if err != nil {
		return nil, fmt.Errorf("failed to register tools: %w", err)
	}

	return ts, nil
}

func (ts *TimeServer) registerTools(localTimezone string) error {
	getCurrentTimeTool := mcp.NewTool("get_current_time",
		mcp.WithDescription("Retrieves current time for a specified timezone"),
		mcp.WithString("timezone",
			mcp.Required(),
			mcp.Description("IANA timezone name (e.g., 'America/New_York', 'Europe/London'). Use 'UTC' as local timezone if no timezone provided by the user."),
		),
	)

	convertTimeTool := mcp.NewTool("convert_time",
		mcp.WithDescription("Convert time between timezones"),
		mcp.WithString("source_timezone",
			mcp.Required(),
			mcp.Description("Source IANA timezone name (e.g., 'America/New_York', 'Europe/London'). Use 'UTC' as local timezone if no source timezone provided by the user."),
		),
		mcp.WithString("time",
			mcp.Required(),
			mcp.Description("Time to convert in 24-hour format (HH:MM)"),
		),
		mcp.WithString("target_timezone",
			mcp.Required(),
			mcp.Description("Target IANA timezone name (e.g., 'Asia/Tokyo', 'America/San_Francisco'). Use 'UTC' as local timezone if no target timezone provided by the user."),
		),
	)

	ts.server.AddTool(getCurrentTimeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		timezone, err := request.RequireString("timezone")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid timezone parameter: %v", err)), nil
		}

		if timezone == "" && localTimezone != "" {
			timezone = localTimezone
		}

		params := types.GetCurrentTimeParams{
			Timezone: timezone,
		}

		result, err := GetCurrentTime(params)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Current time in %s: %s (DST: %t)", result.Timezone, result.Datetime, result.IsDST)), nil
	})

	ts.server.AddTool(convertTimeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		sourceTimezone, err := request.RequireString("source_timezone")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid source_timezone parameter: %v", err)), nil
		}

		time, err := request.RequireString("time")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid time parameter: %v", err)), nil
		}

		targetTimezone, err := request.RequireString("target_timezone")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Invalid target_timezone parameter: %v", err)), nil
		}

		if sourceTimezone == "" && localTimezone != "" {
			sourceTimezone = localTimezone
		}
		if targetTimezone == "" && localTimezone != "" {
			targetTimezone = localTimezone
		}

		params := types.ConvertTimeParams{
			SourceTimezone: sourceTimezone,
			Time:           time,
			TargetTimezone: targetTimezone,
		}

		result, err := ConvertTime(params)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Time conversion:\nSource: %s at %s (%s, DST: %t)\nTarget: %s at %s (%s, DST: %t)\nTime difference: %s",
			result.Source.Timezone, result.Source.Datetime, result.Source.Offset, result.Source.IsDST,
			result.Target.Timezone, result.Target.Datetime, result.Target.Offset, result.Target.IsDST,
			result.TimeDifference)), nil
	})

	return nil
}

func (ts *TimeServer) Serve(ctx context.Context) error {
	return server.ServeStdio(ts.server)
}