FROM golang:1.21.9-alpine AS builder

COPY . /github.com/Arturyus92/auth/source/
WORKDIR /github.com/Arturyus92/auth/source/

RUN go mod download
RUN go build -o ./bin/auth_service cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Arturyus92/auth/source/bin/auth_service .
ADD prod.env .

CMD ["./auth_service"]