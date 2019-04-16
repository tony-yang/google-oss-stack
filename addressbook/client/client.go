package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "addressbook/address"
)

var (
	serverAddress = "127.0.0.1:10000"
)

func printAddressBook(client pb.AddressBookServiceClient) {
	log.Println("Getting AddressBook")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	addressBook, err := client.GetAddressBook(ctx, &pb.AddressBookName{Name: "ABC"})
	if err != nil {
		log.Fatalf("%v.GetAddressBook(_) = _, %v ", client, err)
	}
	log.Println(addressBook)
}

func main() {
	fmt.Println("Starting the Client")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewAddressBookServiceClient(conn)
	printAddressBook(client)
}