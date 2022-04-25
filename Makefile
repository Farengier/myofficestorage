mod:
	@echo "   * Mod  *"
	@go mod tidy
	@go mod vendor

build: mod
	@echo "   * Build *"
	@go build -o bin/storage -mod=vendor src/storage.go

run: build
	@echo "   * Run  *"
	@bin/storage

.PHONY: build run