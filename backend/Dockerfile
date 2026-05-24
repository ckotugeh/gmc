# -------- build --------
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# -------- run --------
FROM alpine:3.20
WORKDIR /app

RUN adduser -D -u 10001 appuser
USER appuser

COPY --from=builder /app/server ./server

EXPOSE 8080

ENV PORT=8080

CMD ["./server"]

