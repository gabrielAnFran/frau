.PHONY: test run
test:
	echo Running GO tests
	mkdir -p cover
	go test -v -failfast -coverprofile "./cover/coverage.out" -coverpkg=./... ./...
	go tool cover -html="./cover/coverage.out" -o ./cover/coverage.html
	go tool cover -func "./cover/coverage.out"
	echo Done
run:
	echo Running docs and app
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init
	go run main.go