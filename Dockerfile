FROM golang:1.24 AS builder

WORKDIR /app

RUN apt-get update && \
    apt-get install -y make && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

COPY . .

RUN go mod download

RUN staticcheck ./...

CMD ["tail", "-f", "/dev/null"]
