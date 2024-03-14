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
	go test -v ./internal/controller/... -run TestIntegration

###################################

# For deployment. Run postgres container also build and run go app container.
deploy:
	docker compose up --build
