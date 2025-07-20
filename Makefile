GOMODNAME := $(shell grep 'module' go.mod | sed -e 's/^module //')
SOURCES := $(shell find . -name "*.go" -or -name "go.mod" -or -name "go.sum" \
	-or -name "Makefile")

# Verbose output
ifdef VERBOSE
V = -v
endif

#
# Environment
#

BINDIR := bin
TOOLDIR := $(BINDIR)/tools

# Global environment variables for all targets
SHELL ?= /bin/bash
SHELL := env \
	GO111MODULE=on \
	GOBIN=$(CURDIR)/$(TOOLDIR) \
	CGO_ENABLED=1 \
	PATH='$(CURDIR)/$(BINDIR):$(CURDIR)/$(TOOLDIR):$(PATH)' \
	$(SHELL)

#
# Defaults
#

# Default target
.DEFAULT_GOAL := test

#
# Tools
#

# external tool
define tool # 1: binary-name, 2: go-import-path
TOOLS += $(TOOLDIR)/$(1)

$(TOOLDIR)/$(1): Makefile
	mkdir -p $(TOOLDIR)
	GOBIN="$(CURDIR)/$(TOOLDIR)" go install "$(2)"
endef

$(eval $(call tool,gofumpt,mvdan.cc/gofumpt@latest))
$(eval $(call tool,goimports,golang.org/x/tools/cmd/goimports@latest))
$(eval $(call tool,golangci-lint,github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest))

.PHONY: tools
tools: $(TOOLS)

#
# Development
#

BENCH ?= .
TESTARGS ?=

.PHONY: clean
clean:
	rm -f $(TOOLS)
	rm -f ./coverage.out ./coverage.html

.PHONY: test
test:
	go test $(V) -count=1 -race $(TESTARGS) ./...

.PHONY: test-deps
test-deps:
	go test all

.PHONY: lint
lint: $(TOOLDIR)/golangci-lint
	golangci-lint $(V) run

.PHONY: format
format: $(TOOLDIR)/goimports $(TOOLDIR)/gofumpt
	goimports -w . && gofumpt -w .

.SILENT: bench
.PHONY: bench
bench:
	go test $(V) -count=1 -bench=$(BENCH) $(TESTARGS) ./...

.PHONY: cov
cov: coverage.out

.PHONY: cov-html
cov-html: coverage.out
	go tool cover -html=./coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: cov-func
cov-func: coverage.out
	go tool cover -func=./coverage.out

coverage.out: $(SOURCES)
	go test $(V) -covermode=count -coverprofile=./coverage.out ./...

.PHONY: deps
deps:
	go mod download

.PHONY: deps-update
deps-update:
	go get -u ./...
	go mod tidy

.PHONY: build
build:
	go build -o $(BINDIR)/$(shell basename $(GOMODNAME)) .

.PHONY: install
install:
	go install .

.SILENT: check-tidy
.PHONY: check-tidy
check-tidy:
	cp go.mod go.mod.tidy-check
	cp go.sum go.sum.tidy-check
	go mod tidy
	( \
		diff go.mod go.mod.tidy-check && \
		diff go.sum go.sum.tidy-check && \
		rm -f go.mod go.sum && \
		mv go.mod.tidy-check go.mod && \
		mv go.sum.tidy-check go.sum \
	) || ( \
		rm -f go.mod go.sum && \
		mv go.mod.tidy-check go.mod && \
		mv go.sum.tidy-check go.sum; \
		exit 1 \
	)

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  test          Run tests with race detection"
	@echo "  test-deps     Run tests for all dependencies"  
	@echo "  lint          Run golangci-lint"
	@echo "  format        Format code with goimports and gofumpt"
	@echo "  bench         Run benchmarks"
	@echo "  cov           Generate coverage report"
	@echo "  cov-html      Generate HTML coverage report"
	@echo "  cov-func      Show coverage by function"
	@echo "  build         Build binary"
	@echo "  install       Install binary"
	@echo "  clean         Clean generated files and tools"
	@echo "  tools         Install development tools"
	@echo "  deps          Download dependencies"
	@echo "  deps-update   Update dependencies"
	@echo "  check-tidy    Check if go.mod and go.sum are tidy"
	@echo "  help          Show this help message"