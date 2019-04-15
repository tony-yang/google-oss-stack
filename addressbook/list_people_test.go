package main

import (
  "bytes"
  "strings"
  "testing"

  pb "addressbook/address"
)

func TestWritePersonWritesPerson(t *testing.T) {
  buf := new(bytes.Buffer)
  p := pb.Person{
    Id: 1234,
    Name: "John Doe",
    Email: "jdoe@abc.com",
    Phones: []*pb.Person_PhoneNumber{
      {Number: "555-4321", Type: pb.Person_HOME},
    },
  }
  writePerson(buf, &p)
  got := buf.String()
  want := `Person ID: 1234
  Name: John Doe
  E-mail address: jdoe@abc.com
  Home phone #: 555-4321
`
  if got != want {
    t.Errorf("writePerson(%s) => \n\t%q, want %q", p.String(), got, want)
  }
}

func TestListPeopleWritesList(t *testing.T) {
  buf := new(bytes.Buffer)
  in := pb.AddressBook{People: []*pb.Person{
      {
        Name: "Jane Doe",
        Id: 102,
      },
      {
        Name: "Jack Doe",
        Id: 201,
        Email: "jack@abc.com",
        Phones: []*pb.Person_PhoneNumber{
          {Number: "555-555-5555", Type: pb.Person_WORK},
        },
      },
  }}
  ListPeople(buf, &in)

  want := strings.Split(`Person ID: 102
  Name: Jane Doe
Person ID: 201
  Name: Jack Doe
  E-mail address: jack@abc.com
  Work phone #: 555-555-5555
`, "\n")
  got := strings.Split(buf.String(), "\n")
  if len(got) != len(want) {
    t.Errorf(
      "listPeople(%s) => \n\t%q has %d lines, want %d",
      in.String(),
      buf.String(),
      len(got),
      len(want))
  }

  for i := 0; i < len(got); i++ {
    if got[i] != want[i] {
      t.Errorf(
        "listPeople(%s) => \n\tline %d %q, want %q",
        in.String(),
        i,
        got[i],
        want[i])
    }
  }
}
