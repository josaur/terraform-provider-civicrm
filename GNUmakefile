# Build configuration
BINARY_NAME=terraform-provider-civicrm
VERSION?=0.1.0
OS_ARCH?=$(shell go env GOOS)_$(shell go env GOARCH)

# Go configuration
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

# Installation directory for local development
# Windows: %APPDATA%\terraform.d\plugins
# Unix: ~/.terraform.d/plugins
ifeq ($(GOOS),windows)
	INSTALL_DIR=$(APPDATA)/terraform.d/plugins/registry.terraform.io/example/civicrm/$(VERSION)/$(OS_ARCH)
else
	INSTALL_DIR=~/.terraform.d/plugins/registry.terraform.io/example/civicrm/$(VERSION)/$(OS_ARCH)
endif

.PHONY: all build install clean test fmt lint docs

all: build

# Build the provider binary
build:
	go build -o $(BINARY_NAME)

# Install the provider locally for testing
install: build
ifeq ($(GOOS),windows)
	@if not exist "$(INSTALL_DIR)" mkdir "$(INSTALL_DIR)"
	copy $(BINARY_NAME) "$(INSTALL_DIR)\$(BINARY_NAME).exe"
else
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
endif

# Clean build artifacts
clean:
ifeq ($(GOOS),windows)
	@if exist $(BINARY_NAME) del $(BINARY_NAME)
	@if exist $(BINARY_NAME).exe del $(BINARY_NAME).exe
else
	rm -f $(BINARY_NAME)
endif

# Run tests
test:
	go test -v ./...

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run ./...

# Generate documentation (requires tfplugindocs)
docs:
	tfplugindocs generate

# Download dependencies
deps:
	go mod download
	go mod tidy

# Update dependencies
update:
	go get -u ./...
	go mod tidy
