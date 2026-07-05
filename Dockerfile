# ------------------------------------------------------------
# Builder
# ------------------------------------------------------------
FROM golang:1.26.3-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

# Workspace
COPY go.work go.work.sum* ./

# Module files
COPY api/go.mod api/go.sum ./api/
COPY pipeline/go.mod pipeline/go.sum ./pipeline/

RUN go work sync

# Copy source
COPY . .

# Build
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build \
        -C api \
        -trimpath \
        -ldflags="-s -w" \
        -o /out/api \
        ./cmd


# ------------------------------------------------------------
# Runtime
# ------------------------------------------------------------
FROM alpine:3.21

RUN apk add --no-cache \
    ca-certificates \
    tzdata

WORKDIR /app

COPY --from=builder /out/api /api

EXPOSE 8080

ENTRYPOINT ["/api"]