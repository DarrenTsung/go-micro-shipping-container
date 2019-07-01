gen:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto

build:
	docker build -t shippy-service-consignment .

run:
	docker run -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		shippy-service-consignment
