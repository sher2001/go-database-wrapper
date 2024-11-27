build: 
	@go build -o bin/myBase cmd/main.go

run: build
	@./bin/myBase

test:
	@go test -v ./...