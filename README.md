# go-user

This project provides Golang services built with Clean Architecture, featuring isolation integration tests using containers, and isolation unit tests with mocks.

# Features

- [x] Clean architecture implementation.
- [x] Isolation integration tests using containers.
- [x] Isolation unit tests with mock support.

# Code structure

```
├── config/             contains the application configuration.
├── internal/
│   ├── app/            contains the application starter.
│   ├── controller/     contains the presentation layer and handles incoming requests from clients.
│   │   ├── http/
│   │   ├── grpc/
│   ├── usecase/        contains the application business layer.
│   ├── repo/           contains the data access layer.
│   ├── entity/         contains the domain model.
├── go.mod
├── go.sum
├── main.go
├── README.md
```

For presentation layer: Some people call it controller (in our case), delivery, transport, or handler.

For business layer: Some people call it usecase (in our case), service, domain, or application.

For data access layer: Some people call it repo/repository (in our case), or persistence.

# Get Started

## Run application

Run postgres container.

```
make compose-up-postgres
```

Run go app.

```
make go-run
```


## Run test

Test can be without the need to run the application.

Run test integration.

```
make go-test-integration
```

Run test unit.

```
make go-test-unit
```

## Run for deployment

For deployment. Run postgres container also build and run go app container.

```
make deploy
```

## Other command

See `Makefile`.
