FROM golang:1.25-bookworm AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /out/main ./cmd/musmgr/


FROM debian:bookworm-slim AS runner

RUN useradd -m -u 1001 musmgr

WORKDIR /app

COPY --from=builder /out/main /app/main

ENV HOME=/home/musmgr
ENV PDFCPU_CONFIG=/tmp/pdfcpu

RUN mkdir -p /tmp/pdfcpu \
    && chown -R musmgr:musmgr /tmp/pdfcpu /app

USER 1001

EXPOSE 4700
EXPOSE 4701

ENTRYPOINT ["/app/main"]
CMD ["--storage-type", "MINIO"]
