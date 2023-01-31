FROM golang:1.19-alpine as build-air

RUN apk update && \
    apk upgrade && \
    apk add git bash --no-cache && \
    go install github.com/cosmtrek/air@latest

FROM golang:1.19-alpine as dev
COPY --from=build-air /go/bin/air /go/bin/air

RUN apk update && \
    apk upgrade && \
    apk add bash git --no-cache

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /app

COPY . .
RUN go mod download

FROM golang:1.19-alpine as prod

WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
