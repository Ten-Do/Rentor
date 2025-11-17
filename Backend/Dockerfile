
# step 1: build the Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /build

RUN --mount=type=cache,target=/var/cache/apk \
    apk add gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# docker ignore may be added to exclude unnecessary files
COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main ./cmd/rentor/main.go

# step 2: create a lightweight image to run the binary
FROM alpine:latest


RUN --mount=type=cache,target=/var/cache/apk \
    apk add sqlite-libs

WORKDIR /rentor

COPY --from=builder /build/main .

# Copy migrations from the builder stage so the runtime can access them
COPY --from=builder /build/migrations ./migrations

# Copy config directory from the builder stage so the runtime can read ./config/local.yaml
COPY --from=builder /build/config ./config

# Create storage directory expected by the app (storage_path: ./storage/storage.db)
RUN mkdir -p /rentor/storage

# Ensure the app knows where to read config from
ENV CONFIG_PATH=/rentor/config/local.yaml

CMD ["./main"]