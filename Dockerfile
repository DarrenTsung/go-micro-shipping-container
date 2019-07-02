FROM golang:1.12 as builder

ENV GO111MODULE=on
WORKDIR /go/src/github.com/DarrenTsung/go-micro-shipping-container
COPY go.mod go.sum ./
RUN go mod download
RUN mkdir -p bin
ARG SUBDIR

COPY . .
RUN go build -o bin/$SUBDIR $SUBDIR/main.go

# ----------------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ARG SUBDIR

COPY --from=builder /go/src/github.com/DarrenTsung/go-micro-shipping-container/bin/$SUBDIR .

CMD ["./$SUBDIR"]
