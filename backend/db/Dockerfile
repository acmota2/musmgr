FROM golang:1.25-bookworm AS builder

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.27.0

FROM debian:bookworm-slim AS runner

# 1. Install certificates and create a non-root system user
RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && useradd -m -u 10001 goose \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# 2. Copy files and transfer ownership to the non-root user
COPY --from=builder --chown=goose:goose /go/bin/goose /usr/local/bin/goose
COPY --from=builder --chown=goose:goose migrations /app/migrations

# 3. Switch to the non-root user
USER 10001

ENTRYPOINT ["/usr/local/bin/goose"]
CMD ["up", "-dir", "migrations"]
