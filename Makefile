# Lint using golangci-lint.
lint:
	golangci-lint run --config .golangci.yml ./...

###################################

# Remove docker image with tag None.
clear-none-docker-images:
	docker images --filter "dangling=true" -q --no-trunc | xargs docker rmi

###################################

# Run postgres container.
compose-up-postgres:
	docker compose up -d go-user-db-postgres

compose-down-postgres:
	docker compose down go-user-db-postgres

# Run go app.
go-run:
	go run main.go -include-migrate

# or run go using air (live reload golang).
air:
	air -c .air.toml

# Run test integration.
go-test-integration:
	go clean -testcache && \
	go test -v ./internal/controller/grpc -run TestIntegration && \
	go test -v ./internal/controller/http -run TestIntegration && \
	go test -cover ./internal/controller/grpc -run TestIntegration && \
	go test -cover ./internal/controller/http -run TestIntegration

# Run test unit.
go-test-unit:
	go clean -testcache && \
	go test -v ./internal/controller/grpc -run TestUnit && \
	go test -v ./internal/controller/http -run TestUnit && \
	go test -v ./internal/usecase -run TestUnit && \
	go test -v ./internal/repo -run TestUnit && \
	go test -cover ./internal/controller/grpc -run TestUnit && \
	go test -cover ./internal/controller/http -run TestUnit && \
	go test -cover ./internal/usecase -run TestUnit && \
	go test -cover ./internal/repo -run TestUnit

# Run test http client.
go-test-http-client:
	go clean -testcache && \
	go test -v ./pkg/gouserhttp -run TestHTTPClient && \
	go test -cover ./pkg/gouserhttp -run TestHTTPClient

# Run test all.
go-test-all:
	make go-test-http-client && make go-test-integration && make go-test-unit 	

###################################

# For deployment. Run postgres container also build and run go app container.
deploy:
	docker compose up --build

###################################

# Generate proto file.
generate-proto:
	protoc \
		--go_out=.      --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/gousergrpc/*.proto
