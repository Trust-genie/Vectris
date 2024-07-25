build:
	@go build -o bin/GoMatrix  ./cmd

run:	build
	@./bin/GoMatrix

test:
	@go test ./... -v