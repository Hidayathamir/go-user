# Lint using golangci-lint.
lint:
	golangci-lint run --config .golangci.yml ./...

###################################

# Remove docker image with tag None.
clear-none-docker-images:
	docker images --filter "dangling=true" -q --no-trunc | xargs docker rmi

###################################

# Run postgres container.
compose-postgres:
	docker compose up go-user-db-postgres

# Run go app.
go-run:
	go run main.go -include-migrate

# Run integration test.
go-integration-test:
	go clean -testcache && \
	go test -v ./integration-test/...

###################################

# For deployment. Run postgres container also build and run go app container.
deploy:
	docker compose up --build

###################################

# Run auth integration test using docker.
docker-auth-integration-test:
	docker \
		compose -f ./integration-test/auth_integration_test_docker_compose.yml \
		up --build --abort-on-container-exit --exit-code-from auth-integration-test-go-user && \
	docker \
		compose -f ./integration-test/auth_integration_test_docker_compose.yml \
		down && \
	docker rmi auth-integration-test-go-user-app auth-integration-test-go-user

docker-profile-integration-test:
	docker \
		compose -f ./integration-test/profile_integration_test_docker_compose.yml \
		up --build --abort-on-container-exit --exit-code-from profile-integration-test-go-user && \
	docker \
		compose -f ./integration-test/profile_integration_test_docker_compose.yml \
		down && \
	docker rmi profile-integration-test-go-user-app profile-integration-test-go-user

docker-integration-test:
	make docker-auth-integration-test && \
	make docker-profile-integration-test && \
	echo "\n\tintegration test success\n"
