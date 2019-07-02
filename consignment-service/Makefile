gen:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment.proto

build:
	docker build -t shippy-service-consignment .

run:
	docker run -p 60051:60051 \
		-e MICRO_SERVER_ADDRESS=:60051 \
		shippy-service-consignment
