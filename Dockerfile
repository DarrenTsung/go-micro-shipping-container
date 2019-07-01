FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shippy-service-consignment

COPY . .

RUN go mod download

# The flags will allow us to run this binary in Alpine.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix ogo -o shippy-service-consignment

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-service-consignment/shippy-service-consignment .

CMD ["./shippy-service-consignment"]
