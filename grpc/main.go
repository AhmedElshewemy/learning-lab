package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc/server/addressbook"
	"grpc/server"

)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Error starting gRPC server: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterAddressBookServiceServer(grpcServer, server.NewAddressbookServer())
	log.Println("gRPC server listening on localhost:50051")
	grpcServer.Serve(listener)
}

