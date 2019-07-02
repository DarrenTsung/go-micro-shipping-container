gen:
	protoc -I. --go_out=plugins=micro:. \
		consignment-service/proto/consignment.proto \
		vessel-service/proto/vessel.proto

build:
	docker build \
		--build-arg SUBDIR=consignment-service \
		-t shippy-service-consignment .
	docker build \
		--build-arg SUBDIR=vessel-service\
		-t shippy-service-vessel .

run-vessel:
	docker run -p 60052:60052 \
		-e MICRO_SERVER_ADDRESS=:60052 \
		shippy-service-vessel

run-consignment:
	docker run -p 60051:60051 \
		-e MICRO_SERVER_ADDRESS=:60051 \
		shippy-service-consignment
