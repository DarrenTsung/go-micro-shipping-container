package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/DarrenTsung/go-micro-shipping-container/consignment-service/proto"
	"github.com/micro/go-micro"
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

	consignment := &pb.Consignment{
		Description: "This is a test consignment",
		Weight:      55000,
		Containers: []*pb.Container{
			&pb.Container{
				CustomerId: "customer001",
				UserId:     "user001",
				Origin:     "Manchester, United Kingdom",
			},
			&pb.Container{
				CustomerId: "customer002",
				UserId:     "user001",
				Origin:     "Manchester, United Kingdom",
			},
			&pb.Container{
				CustomerId: "customer003",
				UserId:     "user002",
				Origin:     "Sheffield, United Kingdom",
			},
		},
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
