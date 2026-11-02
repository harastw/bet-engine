FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o engine cmd/engine/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/engine .
EXPOSE 50051 6060
CMD ["./engine"]
