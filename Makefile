test:
	go test ./... -v

coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

client:
	go run cmd/agent/main.go

server:
	go run cmd/server/main.go

vet:
	go vet ./...

.DEFAULT_GOAL := lint
