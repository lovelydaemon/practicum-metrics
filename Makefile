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

mock:
	mockgen -source ./internal/server/services/interfaces.go -destination ./internal/server/services/mocks_test.go -package services_test

.DEFAULT_GOAL := vet
