# Makefile

# Go command
GO = go

# Output binary name
BINARY = bin/frau

# Build target
.PHONY: build
build:
	$(GO) build -o $(BINARY)
	
test:
	echo Running GO tests
	mkdir -p cover
	go test -v -failfast -coverprofile "./cover/coverage.out" -coverpkg=./... ./...
	go tool cover -html="./cover/coverage.out" -o ./cover/coverage.html
	go tool cover -func "./cover/coverage.out"
	echo Done