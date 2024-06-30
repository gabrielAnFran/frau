# Makefile

# Go command
GO = go

# Output binary name
BINARY = bin/frau

# Build target
.PHONY: build
build:
	$(GO) build -o $(BINARY)

# Clean up build files
.PHONY: clean
clean:
	rm -f $(BINARY)