# --- Stage 1: Build compilation ---
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# FIX: Ensure the source path explicitly points to your backend folder name
COPY backend/go.mod backend/go.sum* ./
RUN go mod download

COPY backend/ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o doctor-platform ./cmd/server


# --- Stage 2: Final runtime container execution ---
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Pull binary from compiler layer stage safely
COPY --from=builder /app/doctor-platform .

# CRITICAL: Recreate local upload storage targets inside the runtime space
RUN mkdir -p uploads/avatars uploads/posts uploads/groups

# Copy base project rules context
COPY backend/.env* ./

EXPOSE 8080

CMD ["./doctor-platform"]
