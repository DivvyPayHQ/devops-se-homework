package main

import (
	"context"
	"flag"
	"log"
	"time"

	// pb "gitlab.com/<gitlab-username>/planet-express/ship/pkg/planetexpress"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func getShip(client pb.PlanetExpressClient) (pb.Ship, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetShip(ctx, &empty.Empty{})

	if err != nil {
		log.Fatalf("%v.GetShip(_) = _, %v: ", client, err)
		return pb.Ship{}, err
	}

	return *resp.Ship, nil
}

func main() {
	log.Println("Planet Express Headquarters")

	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPlanetExpressClient(conn)
	log.Printf("Connected to planet express ship with addr: %s\n\n", *serverAddr)

	ship, _ := getShip(client)
	log.Println(ship)
}
