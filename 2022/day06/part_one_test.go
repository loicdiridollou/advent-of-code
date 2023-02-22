package main

import (
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
  var dat []byte
  dat, _ = os.ReadFile("./day06.testinput")
  actual := part1(string(dat))
  expected := 5
  
  if actual != expected {
    t.Fatal("Error")
  }
}

