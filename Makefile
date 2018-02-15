
GO ?= go

# Build all files.
build: generate
	@echo "==> Build"
	@$(GO) build
.PHONY: build

# Generate all files.
generate:
	@echo "==> Generate"
	@$(GO) generate ./...
.PHONY: generate

# Test all packages.
test: build
	@echo "==> Test"
	@go test -cover ./...
.PHONY: test

# Clean build artifacts, (any non-tracked files in fact, be careful!)
clean:
	@echo "==> Clean"
	@rm -rf dist
#	@git clean -fx
.PHONY: clean

# Release binaries to GitHub.
release: test
	@echo "==> Release"
	@goreleaser --rm-dist
	@echo "==> Complete"
.PHONY: release
