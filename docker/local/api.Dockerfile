FROM golang:1.25-alpine

WORKDIR /app

ENV CGO_ENABLED=1

RUN apk add --no-cache tzdata

RUN apk add --no-cache build-base

CMD ["go", "run", "cmd/api/main.go"]
