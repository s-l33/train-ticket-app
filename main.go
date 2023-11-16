package main

import (
	"github.com/s-l33/train-ticket-app/api"
	"github.com/s-l33/train-ticket-app/api/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	tt := api.NewTrainTicketAppServer()
	pb.RegisterTrainTicketAppServer(s, tt)

	log.Printf("server listening on port %v", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
