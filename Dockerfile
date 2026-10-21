FROM golang:1.24 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    PATH="/go/bin:$PATH"

RUN apt-get update && \
    apt-get install -y make && \
    go install honnef.co/go/tools/cmd/staticcheck@v0.6.1

WORKDIR /app

EXPOSE 3003/tcp
