package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// pb "gitlab.com/<gitlab-username>/planet-express/ship/pkg/planetexpress"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const port = 10000

type planetExpressShipServer struct {
	pb.UnimplementedPlanetExpressServer
}

func newServer() *planetExpressShipServer {
	s := &planetExpressShipServer{}
	return s
}

func (s *planetExpressShipServer) GetShip(ctx context.Context, empty *empty.Empty) (*pb.GetShipResponse, error) {
	return &pb.GetShipResponse{
		Ship: &pb.Ship{
			Name:      "planet express ship",
			Location:  "omicron persei 8",
			FuelLevel: pb.Ship_FULL,
		},
	}, nil
}

func main() {
	log.Println("Planet Express Ship")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterPlanetExpressServer(grpcServer, newServer())

	log.Printf("Ship listening on localhost:%d\n", port)
	grpcServer.Serve(lis)
}
