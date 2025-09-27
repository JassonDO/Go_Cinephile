FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o api cmd/api/main.go
RUN go build -o crawler cmd/crawler/main.go
RUN go build -o scheduler cmd/scheduler/main.go
