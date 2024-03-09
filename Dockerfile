FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o go-user main.go

CMD ["./go-user", "-include-migrate", "-load-env"]
