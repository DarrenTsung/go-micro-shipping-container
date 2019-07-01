gen:
	protoc -I. --go_out=plugins=micro:. \
		proto/consignment/consignment.proto

run:
	go run main.go
