FROM golang:1.12 as builder

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/DarrenTsung/go-micro-shipping-container/cli

COPY . .

RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# ----------------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/DarrenTsung/go-micro-shipping-container/cli/cli .
COPY consignment.json .

CMD ["./cli"]
