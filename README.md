<div align="center">

# go-mcp-time

**A Model Context Protocol (MCP) server for time operations.**

[![Latest Release](https://img.shields.io/github/release/jimeh/go-mcp-time.svg)](https://github.com/jimeh/go-mcp-time/releases)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/jimeh/go-mcp-time)
[![GitHub Issues](https://img.shields.io/github/issues/jimeh/go-mcp-time.svg)](https://github.com/jimeh/go-mcp-time/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/jimeh/go-mcp-time.svg)](https://github.com/jimeh/go-mcp-time/pulls)
[![License](https://img.shields.io/github/license/jimeh/go-mcp-time.svg)](https://github.com/jimeh/go-mcp-time/blob/main/LICENSE)

</div>

## Description

go-mcp-time is a Model Context Protocol (MCP) server that provides time-related operations for AI assistants and other MCP clients. It offers functionality for getting current time in different timezones and converting time between timezones.

The server implements the MCP specification, allowing AI assistants to query time information across various IANA timezone names. It provides a simple and flexible solution for time operations within the MCP ecosystem.

## Installation

```bash
go install github.com/jimeh/go-mcp-time@latest
```

## Usage

### Running the Server

```bash
# Use system default timezone
go-mcp-time

# Override local timezone
go-mcp-time -local-timezone="America/New_York"
```

### Available Tools

The server provides the following MCP tools:

- **get_current_time**: Get current time in a specific timezone
- **convert_time**: Convert time between different timezones

### Example MCP Client Usage

```json
{
  "method": "tools/call",
  "params": {
    "name": "get_current_time",
    "arguments": {
      "timezone": "Europe/London"
    }
  }
}
```

```json
{
  "method": "tools/call",
  "params": {
    "name": "convert_time", 
    "arguments": {
      "time": "14:30",
      "source_timezone": "UTC",
      "target_timezone": "Asia/Tokyo"
    }
  }
}
```

## Documentation

Full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/jimeh/go-mcp-time).

## License

[MIT License](https://github.com/jimeh/go-mcp-time/blob/main/LICENSE)