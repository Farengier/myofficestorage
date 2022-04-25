build:
	@echo "* Build *"
	@go build -o bin/storage src/storage.go
	@echo "* Done  *"

run: build
	@echo "* Run *"
	@bin/storage
	@echo "* Done *"

.PHONY: build run