FROM golang:1.22

WORKDIR /app

COPY . .

CMD ["go", "test", "-v", "./integration-test/profile_integration_test.go"]
