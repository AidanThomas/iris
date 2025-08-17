server:
	@go run cmd/server/main.go
.PHONY: server

client:
	@go run cmd/client/main.go
.PHONY: client
