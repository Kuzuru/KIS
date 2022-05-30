FROM golang:latest as builder

ADD . /go/src/kis/

WORKDIR /go/src/kis/

ENV GO111MODULE=on

COPY go.mod .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
       go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk update \
  && apk add --no-cache

COPY --from=builder /go/src/kis/main /main

ENTRYPOINT [ "/main" ]