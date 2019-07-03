# syntax=docker/dockerfile:experimental
FROM golang:1.12 as builder

ENV GO111MODULE=on
WORKDIR /go/src/github.com/DarrenTsung/go-micro-shipping-container

ARG SUBDIR
RUN mkdir -p bin
ENV GOCACHE=/go/pkg/gocache

COPY . .

RUN --mount=type=cache,id=gocache-v1,target=/go/pkg \
    CGO_ENABLED=0 GOOS=linux \
    go build -o bin/$SUBDIR \
    $SUBDIR/main.go

# ----------------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ARG SUBDIR

COPY --from=builder /go/src/github.com/DarrenTsung/go-micro-shipping-container/bin/$SUBDIR ./app

CMD ["./app"]
