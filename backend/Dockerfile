# --- Stage 1: Build ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o doctor-platform ./cmd/server

# --- Stage 2: Runtime ---
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata wget

WORKDIR /app

COPY --from=builder /app/doctor-platform .

RUN mkdir -p uploads/avatars uploads/posts uploads/groups

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://127.0.0.1:${PORT:-8080}/health || exit 1

CMD ["./doctor-platform"]
