package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc/server/addressbook"
	"grpc/server"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error starting gRPC server: %v", err)
	}
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	
	pb.RegisterAddressBookServiceServer(grpcServer, server.NewAddressbookServer())
	log.Println("gRPC server listening on :50051")
	grpcServer.Serve(listener)
}

