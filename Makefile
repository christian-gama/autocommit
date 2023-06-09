# Name of the output binary
BINARY_NAME=autocommit

# Output directory
OUTPUT_DIR=bin

# Tag for the release
TAG ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")

# Build the application and put the binary in the output directory.
.PHONY: build
build:
	@echo "Building the application..."
	@go build -o $(OUTPUT_DIR)/$(BINARY_NAME) main.go

# Create the output directory if it does not exist.
$(OUTPUT_DIR):
	@mkdir -p $(OUTPUT_DIR)


# Clean the output directory.
.PHONY: clean
clean:
	@echo "Cleaning the output directory..."
	@rm -rf $(OUTPUT_DIR)

# Move the binary to the user's bin directory to make it available in the terminal.
.PHONY: install
install:
	@echo "Installing the binary..."
	@sudo mv $(OUTPUT_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# Build the application for Linux.
.PHONY: build-linux-amd64
build-linux-amd64:
	@echo "Building the application for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME) main.go
	@chmod +x $(OUTPUT_DIR)/$(BINARY_NAME)
	@tar -czvf $(OUTPUT_DIR)/$(BINARY_NAME)-$(TAG)-linux-amd64.tar.gz -C $(OUTPUT_DIR) $(BINARY_NAME)

# Build the application for macOS.
.PHONY: build-macos-amd64
build-macos-amd64:
	@echo "Building the application for macOS..."
	@GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME) main.go
	@chmod +x $(OUTPUT_DIR)/$(BINARY_NAME)
	@tar -czvf $(OUTPUT_DIR)/$(BINARY_NAME)-$(TAG)-darwin-amd64.tar.gz -C $(OUTPUT_DIR) $(BINARY_NAME)

# Build the application for macOS arm64.
.PHONY: build-macos-arm64
build-macos-arm64:
	@echo "Building the application for macOS arm64..."
	@GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME) main.go
	@chmod +x $(OUTPUT_DIR)/$(BINARY_NAME)
	@tar -czvf $(OUTPUT_DIR)/$(BINARY_NAME)-$(TAG)-darwin-arm64.tar.gz -C $(OUTPUT_DIR) $(BINARY_NAME)

# Build the application for Windows.
.PHONY: build-windows-amd64
build-windows-amd64:
	@echo "Building the application for Windows..."
	@GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME).exe main.go
	@chmod +x $(OUTPUT_DIR)/$(BINARY_NAME)
	@zip -j $(OUTPUT_DIR)/$(BINARY_NAME)-$(TAG)-windows-amd64.zip $(OUTPUT_DIR)/$(BINARY_NAME).exe

# Release the application.
.PHONY: release
release: build-linux-amd64 build-macos-amd64 build-macos-arm64 build-windows-amd64
	@echo "Releasing the application..."
	@gh release create $(TAG) $(OUTPUT_DIR)/$(BINARY_NAME)-$(TAG)-linux-amd64.tar.gz $(OUTPUT_DIR)/$(BINARY_NAME)-$(TAG)-darwin-amd64.tar.gz $(OUTPUT_DIR)/$(BINARY_NAME)-$(TAG)-darwin-arm64.tar.gz $(OUTPUT_DIR)/$(BINARY_NAME)-$(TAG)-windows-amd64.zip

