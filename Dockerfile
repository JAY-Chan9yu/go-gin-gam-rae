FROM --platform=linux/amd64 golang:1.20

WORKDIR /app
COPY /. /app

RUN apt-get update && \
    apt-get install unzip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip && \
    unzip -o protoc-3.9.1-linux-x86_64.zip -d /usr/local bin/protoc && \
    unzip -o protoc-3.9.1-linux-x86_64.zip -d /usr/local include/* && \
    rm -rf protoc-3.9.1-linux-x86_64.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go install github.com/cosmtrek/air@latest

RUN go mod download && go mod verify
RUN protoc --go_out=. --go-grpc_out=. proto/service.proto
RUN go build -o main.go

EXPOSE 9000