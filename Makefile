lint:
	golangci-lint run --config .golangci.yml ./...

run:
	go run main.go -include-migrate

compose-up:
	docker compose up --build --force-recreate --no-deps -d && docker compose logs -f

compose-down:
	docker compose down

clear-none-docker-images:
	docker images --filter "dangling=true" -q --no-trunc | xargs docker rmi
