FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o load-balancer ./cmd/load-balancer

CMD ["./load-balancer"]
