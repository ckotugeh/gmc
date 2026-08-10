# --- Stage 1: Build ---

FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Backend Go modules
COPY backend/go.mod backend/go.sum* ./

RUN go mod download

# Backend source
COPY backend/ ./

# Build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o doctor-platform ./cmd/server


# --- Stage 2: Runtime ---

FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/doctor-platform .

RUN mkdir -p uploads/avatars uploads/posts uploads/groups

EXPOSE 8080

CMD ["./doctor-platform"]