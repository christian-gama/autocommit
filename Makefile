# Name of the output binary
BINARY_NAME=autocommit

# Output directory
OUTPUT_DIR=bin

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
install: build
	@echo "Installing the binary..."
	@sudo mv $(OUTPUT_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)