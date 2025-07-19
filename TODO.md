# Go Time MCP Server - Implementation Plan

## Project Overview

Create a memory and CPU efficient Time MCP server in Go as a drop-in replacement for the Python-based variant. The goal is to build a minimal, fast, and lightweight MCP server with a tiny Docker container image.

## Analysis of Existing Python Implementation

### MCP Tools to Implement

#### 1. `get_current_time`
- **Functionality**: Retrieves current time for a specified timezone
- **Parameters**:
  - `timezone` (required): IANA timezone name (e.g., "Europe/Warsaw")
- **Returns**:
  - `timezone`: The requested timezone
  - `datetime`: Current timestamp with timezone offset
  - `is_dst`: Whether daylight saving time is in effect

#### 2. `convert_time`
- **Functionality**: Converts time between different timezones
- **Parameters**:
  - `source_timezone`: Source IANA timezone name
  - `time`: Time in 24-hour format (HH:MM)
  - `target_timezone`: Target IANA timezone name
- **Returns**:
  - `source`: Original timezone details
  - `target`: Converted timezone details
  - `time_difference`: Time offset between zones

### Key Features from Python Implementation
- Automatic system timezone detection by default
- Support for manual timezone override with `--local-timezone` argument
- IANA timezone names for precise conversions
- Detailed datetime information including offset and DST status
- Error handling for invalid timezones and time formats

## Go Package Selection

### MCP Protocol Library
**Primary Choice**: `github.com/metoro-io/mcp-golang`
- **Rationale**: Unofficial but well-designed, type-safe implementation
- **Features**: 
  - Low boilerplate with automatic endpoint generation
  - Type-safe tool arguments using Go structs
  - Multiple transport options (stdio, HTTP)
  - Modular architecture
  - Active community support

**Alternative**: `github.com/mark3labs/mcp-go`
- **Rationale**: Another solid option, handles protocol details well
- **Note**: Consider if mcp-golang doesn't meet needs

### Testing Framework
**Choice**: `github.com/stretchr/testify/assert`
- **Rationale**: Minimal, widely-used testing assertions
- **Features**: Simple assertion methods, clear error messages

### Additional Standard Library Packages
- `time`: Core time functionality and timezone handling
- `encoding/json`: JSON marshaling/unmarshaling
- `fmt`: String formatting and error messages
- `os`: Environment variable access and CLI arguments

## Implementation Plan

### Phase 1: Project Structure Setup
1. Initialize Go module with `go mod init`
2. Create basic directory structure:
   ```
   ├── main.go
   ├── internal/
   │   ├── server/
   │   │   ├── server.go
   │   │   └── handlers.go
   │   └── types/
   │       └── time.go
   ├── go.mod
   ├── go.sum
   ├── Dockerfile
   └── README.md
   ```

### Phase 2: Core Implementation
1. **Define data structures** (`internal/types/time.go`):
   - `TimeResult` struct for `get_current_time` response
   - `TimeConversionResult` struct for `convert_time` response
   - Input parameter structs for type safety

2. **Implement MCP server** (`internal/server/server.go`):
   - Initialize mcp-golang server with stdio transport
   - Register MCP tools with proper metadata
   - Handle server lifecycle and graceful shutdown

3. **Implement time handlers** (`internal/server/handlers.go`):
   - `GetCurrentTime()` function with timezone validation
   - `ConvertTime()` function with time parsing and conversion
   - Error handling for invalid inputs
   - Timezone validation using Go's time package

4. **Main entry point** (`main.go`):
   - Command-line argument parsing (local timezone override)
   - Server initialization and startup
   - Graceful error handling and logging

### Phase 3: Docker Containerization
1. **Multi-stage Dockerfile**:
   - Stage 1: Build environment with Go toolchain
   - Stage 2: Minimal runtime using `scratch` base image
   - Static linking for dependency-free binary

2. **Build optimizations**:
   - CGO disabled for static linking
   - Strip debug symbols for smaller binary
   - Use Go's built-in timezone data

### Phase 4: Testing & Validation
1. **Unit tests**:
   - Test timezone validation
   - Test time conversion logic
   - Test error handling scenarios
   - Mock MCP protocol interactions

2. **Integration tests**:
   - Test full MCP tool workflows
   - Verify JSON-RPC protocol compliance
   - Test CLI argument handling

3. **Performance testing**:
   - Memory usage benchmarks
   - Response time measurements
   - Comparison with Python implementation

### Phase 5: Documentation & Distribution
1. **Documentation**:
   - Usage examples
   - Configuration options
   - Docker deployment instructions

2. **Build artifacts**:
   - Cross-platform binaries
   - Docker image publishing
   - Release automation

## Docker Strategy

### Dockerfile Structure
```dockerfile
# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o time-server main.go

# Runtime stage
FROM scratch
COPY --from=builder /app/time-server /time-server
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENTRYPOINT ["/time-server"]
```

### Optimization Goals
- **Image size**: Target under 10MB
- **Memory usage**: Under 20MB runtime
- **Startup time**: Sub-second initialization
- **Security**: Distroless/scratch base for minimal attack surface

## Success Criteria

### Functional Requirements
- ✅ Drop-in replacement for Python Time MCP server
- ✅ Support both `get_current_time` and `convert_time` tools
- ✅ Full IANA timezone support
- ✅ CLI compatibility with original server

### Performance Requirements
- ✅ 50%+ reduction in memory usage vs Python
- ✅ 2x+ faster response times
- ✅ 5x+ smaller container image
- ✅ Sub-second startup time

### Quality Requirements
- ✅ 90%+ test coverage
- ✅ No runtime dependencies
- ✅ Cross-platform compatibility
- ✅ Production-ready error handling

## Next Steps

1. Initialize Go module and basic project structure
2. Implement core MCP server with mcp-golang
3. Add time conversion functionality with proper timezone handling
4. Create comprehensive test suite
5. Build optimized Docker container
6. Performance benchmarking and optimization
7. Documentation and deployment guides