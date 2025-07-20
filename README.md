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

### Using Go Binary

Start the MCP server using the command line:

```bash
# Use system default timezone
go-mcp-time

# Override local timezone (accepts IANA timezone names)
go-mcp-time -local-timezone="America/New_York"
go-mcp-time -local-timezone="Europe/London"
go-mcp-time -local-timezone="Asia/Tokyo"
```

### Using Docker

A container image is available at `ghcr.io/jimeh/go-mcp-time`. The image is built from scratch and includes timezone data for IANA timezone support. When running via Docker, the server defaults to UTC timezone.

```bash
# Use default UTC timezone
docker run --rm ghcr.io/jimeh/go-mcp-time

# Set timezone using TZ environment variable
docker run --rm -e TZ="America/New_York" ghcr.io/jimeh/go-mcp-time
docker run --rm -e TZ="Europe/London" ghcr.io/jimeh/go-mcp-time

# Run with specific tag/version
docker run --rm -e TZ="Asia/Tokyo" ghcr.io/jimeh/go-mcp-time:v0.0.2
```

The container exposes the MCP server on stdout/stdin for communication with MCP clients.

### Command Line Options

- `-local-timezone`: Override the local timezone (default: system timezone or UTC)

### Available Tools

The server provides the following MCP tools:

- **get_current_time**: Get current time in a specific timezone
- **convert_time**: Convert time between different timezones

## Documentation

Full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/jimeh/go-mcp-time).

## License

[MIT License](https://github.com/jimeh/go-mcp-time/blob/main/LICENSE)