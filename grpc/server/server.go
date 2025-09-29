package server

import (
    "context"
    "log"

    pb "grpc/server/addressbook"
)

type AddressbookServer struct {
    pb.UnimplementedAddressBookServiceServer
}

func NewAddressbookServer() *AddressbookServer {
    return &AddressbookServer{}
}

func (s *AddressbookServer) AddPerson(ctx context.Context, person *pb.Person) (*pb.AddPersonResponse, error) {
    log.Println("Received:", person)
    return &pb.AddPersonResponse{}, nil
}

func (s *AddressbookServer) ListPersons(ctx context.Context, req *pb.ListPersonsRequest) (*pb.ListPersonsResponse, error) {

    log.Println("ListPersons called")

    return &pb.ListPersonsResponse{}, nil

}
