FROM golang:1.12

RUN apt-get update && apt-get install -y unzip
WORKDIR /tmp
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip
RUN unzip protoc-3.8.0-linux-x86_64.zip
RUN mv ./bin/protoc /usr/local/bin
RUN go get -u github.com/micro/protobuf/proto github.com/micro/protobuf/protoc-gen-go

WORKDIR /app

CMD ["bash"]
