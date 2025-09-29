package main
import (
	"log"
	"context"

	"google.golang.org/grpc"
	pb "grpc/server/addressbook"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAddressBookServiceClient(conn)

	// Call the AddPerson method

	resp, err := client.AddPerson(context.Background(), &pb.Person{Name: "John Doe", Email: "jle.com"})
	if err != nil {
		log.Fatalf("could not add person: %v", err)
	}
	log.Printf("AddPerson response: %v", resp)


}