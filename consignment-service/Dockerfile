FROM golang:1.12 as builder

WORKDIR /go/src/github.com/DarrenTsung/go-micro-shipping-container/consignment-service

RUN go get -u github.com/golang/dep/cmd/dep

COPY . .

RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# ----------------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/DarrenTsung/go-micro-shipping-container/consignment-service/consignment-service .

CMD ["./consignment-service"]
