FROM golang:1.25-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main ./cmd/musmgr/
EXPOSE 4700
EXPOSE 4701

ENTRYPOINT ["/app/main"]
CMD ["--storage-type", "MINIO"]
