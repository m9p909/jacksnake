

test: 
	go test ./...

run: 
	go run .

build:
	go build .

devtest:
	gow test ./...

devrun:
	gow run .

ci: test build
