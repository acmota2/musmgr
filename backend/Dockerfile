FROM golang:1.21.5
WORKDIR /usr/src/app
COPY . .
RUN go build -o backend main.go
EXPOSE 3000
CMD ["./backend"]
