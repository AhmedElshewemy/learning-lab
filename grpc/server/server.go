package server

import (
    "context"
    "fmt"

    pb "grpc/server/addressbook"
)

type AddressbookServer struct {
    pb.UnimplementedAddressBookServiceServer
}

func NewAddressbookServer() *AddressbookServer {
    return &AddressbookServer{}
}

func (s *AddressbookServer) AddPerson(ctx context.Context, person *pb.Person) (*pb.Person, error) {
    fmt.Println("Received:", person)
    return person, nil
}
