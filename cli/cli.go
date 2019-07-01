package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/DarrenTsung/go-micro-shipping-container/proto/consignment"
	"github.com/micro/go-micro"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()

	client := pb.NewShippingServiceClient("shippy.consignment.service", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not create consignment: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	listResponse, err := client.ListConsignments(context.Background(), &pb.ListRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range listResponse.Consignments {
		log.Println(v)
	}
}
