FROM golang:latest as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o main ./cmd/api/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/
COPY ./cmd/api/app.env /app/app.env
ENTRYPOINT ["./main"]