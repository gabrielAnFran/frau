.PHONY: test run
test:
	echo Running GO tests
	mkdir -p cover
	go test -v -failfast -coverprofile "./cover/coverage.out" -coverpkg=./... ./...
	go tool cover -html="./cover/coverage.out" -o ./cover/coverage.html
	go tool cover -func "./cover/coverage.out"

build:
	echo Building FrauClI
	go build -o "./bin/frau"
	
