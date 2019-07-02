package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/DarrenTsung/go-micro-shipping-container/vessel-service/proto"
	"github.com/micro/go-micro"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

type service struct {
	repo repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.FindResponse) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatyface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(micro.Name("shippy.service.vessel"))
	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("Server finished, exiting normally")
}
