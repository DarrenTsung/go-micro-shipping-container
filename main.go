package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/DarrenTsung/go-micro-shipping-container/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	List() []*pb.Consignment
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

func (repo *Repository) List() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.CreateResponse, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{Created: true, Consignment: consignment}, nil
}

func (s *service) ListConsignments(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	return &pb.ListResponse{Consignments: s.repo.List()}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
	log.Println("Server finished, exiting normally")
}
