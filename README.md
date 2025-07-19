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

### Configuration Examples

#### Claude Code

Add to your MCP settings in `~/.config/claude-code/mcp_servers.json`:

```json
{
  "mcpServers": {
    "go-mcp-time": {
      "command": "go-mcp-time",
      "args": ["-local-timezone=UTC"]
    }
  }
}
```

#### Claude Desktop

Add to your configuration file:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`  
**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "go-mcp-time": {
      "command": "go-mcp-time",
      "args": ["-local-timezone=UTC"]
    }
  }
}
```

#### Cursor

Add to your MCP configuration in `.cursorrules` or settings:

```json
{
  "mcp": {
    "servers": {
      "go-mcp-time": {
        "command": "go-mcp-time",
        "args": ["-local-timezone=UTC"]
      }
    }
  }
}
```

#### VSCode (with MCP extension)

Add to your `settings.json`:

```json
{
  "mcp.servers": {
    "go-mcp-time": {
      "command": "go-mcp-time",
      "args": ["-local-timezone=UTC"]
    }
  }
}
```

## Documentation

Full documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/jimeh/go-mcp-time).

## License

[MIT License](https://github.com/jimeh/go-mcp-time/blob/main/LICENSE)