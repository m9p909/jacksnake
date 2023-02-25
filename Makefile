

test: 
	go test ./...
test-nocache:
	go test -count=1 ./...
run: 
	go run .

build:
	go build .

devtest:
	gow test ./...

devrun:
	gow run .

ci: test build
