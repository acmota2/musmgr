FROM golang:1.25-bookworm AS builder

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.27.0

FROM debian:bookworm-slim AS runner

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && useradd -m -u 10001 goose \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder --chown=goose:goose /go/bin/goose /usr/local/bin/goose
COPY --chown=goose:goose migrations /app/migrations

USER 10001

ENTRYPOINT ["/usr/local/bin/goose"]
CMD ["up", "-dir", "migrations"]
