
GO ?= go

# Build all files.
build: generate
	@echo "==> Building"
	@$(GO) build
.PHONY: build

# Generate all files.
generate:
	@echo "==> Generating"
	@$(GO) generate ./...
.PHONY: generate

# Test all packages.
test: build
	@echo "==> Testing"
	@go test -cover ./...
.PHONY: test

# Release binaries to GitHub.
release: test
	@echo "==> Releasing"
#	@goreleaser -p 1 --rm-dist -config .goreleaser.yml
	@echo "==> Complete"
.PHONY: release

# Clean build artifacts, (any non-tracked files in fact, be careful!)
clean:
	@echo "==> Cleaning"
	@git clean -fx
.PHONY: clean
