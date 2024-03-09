lint:
	golangci-lint run --config .golangci.yml ./...

migrate:
	go run main.go -migrate

run:
	go run main.go
