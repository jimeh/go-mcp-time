// Package main implements a Go Time MCP server for time conversion operations.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jimeh/go-mcp-time/server"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	var localTimezone string
	var showVersion bool
	flag.StringVar(&localTimezone, "local-timezone", "",
		"Override local timezone (IANA timezone name)")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.Parse()

	if showVersion {
		fmt.Printf("go-mcp-time %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("built: %s\n", date)

		os.Exit(0)
	}

	if localTimezone == "" {
		// Check TZ environment variable first
		if tz := os.Getenv("TZ"); tz != "" {
			localTimezone = tz
		} else {
			// Try to get the system timezone location
			now := time.Now()
			if loc := now.Location(); loc != nil && loc.String() != "" {
				localTimezone = loc.String()
			} else {
				localTimezone = "UTC"
			}
		}
	}

	timeServer, err := server.NewTimeServer(localTimezone)
	if err != nil {
		log.Fatalf("Failed to create time server: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Fprintf(os.Stderr,
			"\nReceived interrupt signal, shutting down...\n")
		cancel()
	}()

	if err := timeServer.Serve(ctx); err != nil {
		log.Printf("Server error: %v", err)
	}
}
