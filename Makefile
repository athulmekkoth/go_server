build:
	@go build -o bin/server_go cmd/migrate/main.go

run: build
	@./bin/server_go
