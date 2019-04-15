package main

import (
  "strings"
  "testing"
)

func TestPromptForAddressReturnsAddress(t *testing.T) {
  in := `12345
John Lark
lark@abc.com
123-456-7890
home
222-222-2222
mobile
777-777-7777
unknown

`
  got, err := PromptForAddress(strings.NewReader(in))
  if err != nil {
    t.Fatalf("PromptForAddress(%q) had unexpected error: %s", in, err.Error())
  }
  if got.Id != 12345 {
    t.Errorf("promptForAddress(%q) got %d, want ID %d", in, got.Id, 12345)
  }
  if got.Name != "John Lark" {
    t.Errorf("prmptForAddress(%q) => want name %q, got %q", in, "John Lark", got.Name)
  }
  if got.Email != "lark@abc.com" {
    t.Errorf("prmptForAddress(%q) => want email %q, got %q", in, "lark@abc.com", got.Email)
  }
}
