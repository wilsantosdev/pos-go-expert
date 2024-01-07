FROM go:latest as builder
WORKDIR /app
COPY ./../../ /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a cgo -o main ./cmd/api/main.go

