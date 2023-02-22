package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
  var dat []byte
  dat, _ = os.ReadFile("./day02.testinput")
  actual := part1(string(dat))
  expected := 15
  
  if actual != expected {
    t.Fatal("Error")
  }
}

