FROM golang:1.25 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# docker ignore
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/rentor/main.go

FROM alpine:latest

WORKDIR /rentor

COPY --from=builder /build/main .

# Copy config directory from the builder stage so the runtime can read ./config/local.yaml
COPY --from=builder /build/config ./config

# Create storage directory expected by the app (storage_path: ./storage/storage.db)
RUN mkdir -p /rentor/storage

# Ensure the app knows where to read config from
ENV CONFIG_PATH=/rentor/config/local.yaml

CMD ["./main"]