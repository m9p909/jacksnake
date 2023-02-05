

test: 
	go test ./...

run: 
	go run .

build:
	go build .

ci: test build
