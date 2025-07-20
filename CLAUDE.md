# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

go-mcp-time is a Go implementation of a Model Context Protocol (MCP) server that provides time-related operations. It offers two main tools: `get_current_time` for retrieving current time in different timezones, and `convert_time` for converting time between timezones using IANA timezone names.

## Development Commands

### Essential Commands
- `make test` - Run tests with race detection (default target)
- `make lint` - Run golangci-lint for code quality checks
- `make format` - Format code with goimports and gofumpt
- `make build` - Build the binary to `bin/go-mcp-time`
- `make install` - Install the binary globally

### Testing and Coverage
- `make cov` - Generate coverage report (`coverage.out`)
- `make cov-html` - Generate HTML coverage report (`coverage.html`)
- `make cov-func` - Show coverage by function
- `make bench` - Run benchmarks

### Development Tools
- `make tools` - Install development tools (gofumpt, goimports, golangci-lint)
- `make deps` - Download dependencies
- `make deps-update` - Update all dependencies
- `make check-tidy` - Verify go.mod and go.sum are tidy

## Code Architecture

### Package Structure
- **main.go** - CLI entry point with flag parsing and signal handling
- **server/** - Core MCP server implementation
  - `server.go` - MCP server setup and tool registration using mark3labs/mcp-go
  - `handlers.go` - Time operation handlers with timezone validation
- **types/** - Data structures for time operations
  - `time.go` - Request/response types for MCP tools

### Key Design Principles
- **Importable packages**: All logic in named packages (server, types) for reusability
- **MCP compliance**: Uses mark3labs/mcp-go library for protocol implementation
- **Timezone handling**: IANA timezone names with Go's time.LoadLocation()
- **Error handling**: Comprehensive validation for timezones and time formats
- **Testing**: Comprehensive unit tests with testify/assert

### MCP Tools Implementation
1. **get_current_time**: Uses `time.Now().In(loc)` for timezone-aware current time
2. **convert_time**: Validates HH:MM format with regex, converts using `time.Date()` and `In()`

### Dependencies
- `github.com/mark3labs/mcp-go` - MCP protocol implementation
- `github.com/stretchr/testify` - Testing assertions
- Standard library: `time`, `regexp`, `context`, `flag`

### Running the Server
```bash
# Use system timezone
go-mcp-time

# Override local timezone
go-mcp-time -local-timezone="America/New_York"
```

## Release Management

This project uses **release-please** for automated release and changelog management. This requires strict adherence to conventional commit format.

### Commit Message Format
All commits and PR titles MUST follow [Conventional Commits](https://www.conventionalcommits.org/) format:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Commit Types
- `feat` - New features
- `fix` - Bug fixes  
- `docs` - Documentation changes
- `test` - Test additions/modifications
- `refactor` - Code refactoring
- `ci` - CI/CD changes
- `chore` - Maintenance tasks

### Recommended Scopes
- `server` - Changes to server package
- `types` - Changes to types package
- `cli` - Changes to main.go/CLI functionality
- `deps` - Dependency updates
- `docker` - Docker/containerization changes

### Examples
- `feat(server): add timezone validation for MCP tools`
- `fix(types): correct JSON serialization for TimeResult`
- `docs(readme): update installation instructions`
- `test(server): add comprehensive timezone conversion tests`
- `chore(deps): update mark3labs/mcp-go to v0.34.0`

### Test Patterns
- Table-driven tests for different timezone scenarios
- Regex validation for datetime format outputs
- Error message assertion for invalid inputs
- DST and offset validation for timezone conversions