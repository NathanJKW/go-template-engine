build:
	go build -o ./bin/cli ./cmd/cli/main.go

test:
	go test ./...