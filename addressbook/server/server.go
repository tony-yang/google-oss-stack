package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "addressbook/address"
)

var (
	port = 10000
)

type addressBookServer struct {}

func (s *addressBookServer)GetAddressBook(ctx context.Context, filename *pb.AddressBookName) (*pb.AddressBook, error) {
	book := &pb.AddressBook{}
	// Hard coding the return data for demo purpose
	addr := &pb.Person{
		Name: "John Doe",
		Id: 1234,
	}
	addr2 := &pb.Person{
		Name: "Jane Doe",
		Id: 1235,
		Email: "jane@abc.com",
	}
	book.People = append(book.People, addr)
	book.People = append(book.People, addr2)
	return book, nil
}

func newServer() *addressBookServer {
	s := &addressBookServer{}
	return s
}

func main() {
	fmt.Println("Starting the Server")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAddressBookServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}