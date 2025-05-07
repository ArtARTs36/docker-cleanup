# syntax=docker/dockerfile:1

FROM golang:1.23.0-alpine AS builder

ARG APP_VERSION="undefined"
ARG BUILD_TIME="undefined"

WORKDIR /go/src/github.com/artarts36/docker-cleanup

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w -X 'main.Version=${APP_VERSION}' -X 'main.BuildDate=${BUILD_TIME}'" -o /go/bin/docker-cleanup /go/src/github.com/artarts36/docker-cleanup/cmd/docker-cleanup/main.go

######################################################

FROM scratch

COPY --from=builder /go/bin/docker-cleanup /go/bin/docker-cleanup

EXPOSE 8000

CMD ["/go/bin/docker-cleanup"]
