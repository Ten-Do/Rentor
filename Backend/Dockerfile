FROM golang:1.25 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# почитать docker ignore
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/rentor/main.go

FROM alpine:latest

WORKDIR /rentor

COPY --from=builder /build/main .

CMD ["./main"]