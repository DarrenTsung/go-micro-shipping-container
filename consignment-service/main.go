package main

import (
	"context"
	"log"
	"sync"

	pb "github.com/DarrenTsung/go-micro-shipping-container/consignment-service/proto"
	"github.com/micro/go-micro"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	ListAll() []*pb.Consignment
}

// Dummy repository
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated

	return consignment, nil
}

func (repo *Repository) ListAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.CreateResponse) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment

	return nil
}

func (s *service) ListConsignments(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	res.Consignments = s.repo.ListAll()
	return nil
}

func main() {
	repo := &Repository{}
	srv := micro.NewService(
		micro.Name("shippy.consignment.service"),
	)

	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("Server finished, exiting normally")
}
