FROM golang:latest as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o main main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main
ENTRYPOINT ["./main"]