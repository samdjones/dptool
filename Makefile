
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

# Build a distribution (using GoReleaser for convenience).
dist: test
	@echo "==> Dist"
	@goreleaser --rm-dist --snapshot
	@echo "==> Complete"
.PHONY: dist

# Release binaries to GitHub.
release: test
	@echo "==> Release"
#	@goreleaser -p 1 --rm-dist -config .goreleaser.yml
	@echo "==> Complete"
.PHONY: release
