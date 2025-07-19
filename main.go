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

	"github.com/jimeh/go-time-mcp/server"
)

func main() {
	var localTimezone string
	flag.StringVar(&localTimezone, "local-timezone", "", "Override local timezone (IANA timezone name)")
	flag.Parse()

	if localTimezone == "" {
		zone, _ := time.Now().Zone()
		if zone != "" {
			localTimezone = zone
		} else {
			localTimezone = "UTC"
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
		fmt.Fprintf(os.Stderr, "\nReceived interrupt signal, shutting down...\n")
		cancel()
	}()

	if err := timeServer.Serve(ctx); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}