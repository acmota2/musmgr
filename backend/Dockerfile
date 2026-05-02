FROM golang:1.25-bookworm

RUN useradd -m -u 1001 appuser

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main ./cmd/musmgr/

RUN chown appuser:appuser /app/main

ENV PDFCPU_CONFIG=/tmp/pdfcpu

USER 1001

EXPOSE 4700
EXPOSE 4701

ENTRYPOINT ["/app/main"]
CMD ["--storage-type", "MINIO"]
