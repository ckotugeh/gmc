# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (required by some Go modules)
RUN apk add --no-cache git

# Copy Go module files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source
COPY . .

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o doctor-platform ./cmd/server

# ---------- Runtime Stage ----------
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the binary
COPY --from=builder /app/doctor-platform .

# Copy environment file if present
COPY .env* ./

EXPOSE 8080

CMD ["./doctor-platform"]