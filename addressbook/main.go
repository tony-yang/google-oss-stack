package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "strings"

  "github.com/golang/protobuf/proto"
  pb "addressbook/address"
)

func main() {
  if len(os.Args) != 3 {
    log.Fatalf("Usage: %s ACTION ADDRESS_BOOK_FILE\n\n\tACTION: add, list\n", os.Args[0])
  }
  action := strings.ToLower(os.Args[1])
  filename := os.Args[2]

  if action == "add" {
    in, err := ioutil.ReadFile(filename)
    if err != nil {
      if os.IsNotExist(err) {
        fmt.Printf("%s: File not found. Creating new file.\n", filename)
      } else {
        log.Fatalln("Error reading file:", err)
      }
    }

    book := &pb.AddressBook{}
    if err := proto.Unmarshal(in, book); err != nil {
      log.Fatalln("Failed to parse address book:", err)
    }

    addr, err := PromptForAddress(os.Stdin)
    if err != nil {
      log.Fatalln("Error with address:", err)
    }
    book.People = append(book.People, addr)

    out, err := proto.Marshal(book)
    if err != nil {
      log.Fatalln("Failed to encode address book:", err)
    }

    if err := ioutil.WriteFile(filename, out, 0644); err != nil {
      log.Fatalln("Failed to write address book:", err)
    }
  } else if action == "list" {
    in, err := ioutil.ReadFile(filename)
    if err != nil {
      log.Fatalln("Erorr reading file:", err)
    }

    book := &pb.AddressBook{}
    if err := proto.Unmarshal(in, book); err != nil {
      log.Fatalln("Failed to parse address book:", err)
    }
    ListPeople(os.Stdout, book)
  } else {
    fmt.Println("Warning: Unknown action")
    log.Fatalf("Usage: %s ACTION ADDRESS_BOOK_FILE\n\n\tACTION: add, list\n", os.Args[0])
  }
}
